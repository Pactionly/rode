package server

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/golang/protobuf/proto"
	"github.com/rode/grafeas-elasticsearch/go/v1beta1/storage"
	pb "github.com/rode/rode/proto/v1alpha1"
	grafeas_proto "github.com/rode/rode/protodeps/grafeas/proto/v1beta1/grafeas_go_proto"
	"go.uber.org/zap"
	"google.golang.org/protobuf/encoding/protojson"
)

func (r *rodeServer) ListResources(ctx context.Context, request *pb.ListResourcesRequest) (*pb.ListResourcesResponse, error) {
	log := r.logger.Named("ListResources")
	log.Debug("received request", zap.Any("ListResourcesRequest", request))

	searchQuery := storage.EsSearch{
		Collapse: &storage.EsCollapse{
			Field: "resource.uri",
		},
	}

	if request.Filter != "" {
		parsedQuery, err := r.filterer.ParseExpression(request.Filter)
		if err != nil {
			log.Error("failed to parse query", zap.Error(err))
			return nil, err
		}

		searchQuery.Query = parsedQuery
	}

	encodedBody, requestJSON := encodeRequest(searchQuery)
	log.Debug("es request payload", zap.Any("payload", requestJSON))
	res, err := r.esClient.Search(
		r.esClient.Search.WithContext(ctx),
		r.esClient.Search.WithIndex("grafeas-v1beta1-rode-occurrences"),
		r.esClient.Search.WithBody(encodedBody),
		r.esClient.Search.WithSize(1000),
	)

	if err != nil {
		return nil, err
	}

	if res.IsError() {
		return nil, fmt.Errorf("error occurred during ES query %v", res)
	}

	var searchResults storage.EsSearchResponse
	if err := decodeResponse(res.Body, &searchResults); err != nil {
		return nil, err
	}
	var resources []*grafeas_proto.Resource
	for _, hit := range searchResults.Hits.Hits {
		hitLogger := log.With(zap.String("raw occurrence", string(hit.Source)))

		occurrence := &grafeas_proto.Occurrence{}
		err := protojson.Unmarshal(hit.Source, proto.MessageV2(occurrence))
		if err != nil {
			log.Error("failed to convert", zap.Error(err))
			return nil, err
		}

		hitLogger.Debug("occurrence hit", zap.Any("occurrence", occurrence))

		resources = append(resources, occurrence.Resource)
	}

	return &pb.ListResourcesResponse{
		Resources:     resources,
		NextPageToken: "",
	}, nil
}

func encodeRequest(body interface{}) (io.Reader, string) {
	b, err := json.Marshal(body)
	if err != nil {
		// we should know that `body` is a serializable struct before invoking `encodeRequest`
		panic(err)
	}

	return bytes.NewReader(b), string(b)
}

func decodeResponse(r io.ReadCloser, i interface{}) error {
	return json.NewDecoder(r).Decode(i)
}
