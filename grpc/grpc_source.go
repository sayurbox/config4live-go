package grpc

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/afex/hystrix-go/hystrix"
	pb "github.com/sayurbox/config4live-go/config4live"
	"github.com/sayurbox/config4live-go/internal"
	"google.golang.org/grpc"
)

// GrpcSource grpc configuration source
type GrpcSource struct {
	url          string
	stub         pb.LiveConfigurationClient
	hystrixParam *internal.HystrixParams
}

// Get call grpc findConfig by key
func (g *GrpcSource) Get(key string) (*internal.Config, error) {
	hystrix.ConfigureCommand(g.hystrixParam.Name, g.hystrixParam.HystrixConfig())
	responseChannel := make(chan internal.GetConfigResponse, 1)
	hystrix.Go(g.hystrixParam.Name, func() error {
		response, e := g.stub.FindConfig(context.Background(), &pb.ConfigRequest{Name: key})
		responseChannel <- internal.GetConfigResponse{
			Output: parseResponse(response),
			Error:  e,
		}
		return e
	}, func(e error) error {
		log.Println("Fallback is executed")
		responseChannel <- internal.GetConfigResponse{
			Error: e,
		}
		return e
	})
	hystrixResp := <-responseChannel
	if hystrixResp.Error != nil {
		log.Printf("Error get config :%s\n", hystrixResp.Error.Error())
	}

	if hystrixResp.Output != nil && hystrixResp.Output.ID == "" {
		return nil, fmt.Errorf("config %s is not found", key)
	}

	return hystrixResp.Output, hystrixResp.Error
}

// NewGrpcSource create configuration source from grpc
func NewGrpcSource(opts ...Option) *GrpcSource {
	s := &GrpcSource{}
	for _, opt := range opts {
		opt(s)
	}
	if s.hystrixParam == nil {
		s.hystrixParam = &internal.HystrixParams{}
	}
	if s.hystrixParam.Name == "" {
		s.hystrixParam.Name = "live-config-command-key"
	}
	conn, err := grpc.Dial(s.url, grpc.WithInsecure(), grpc.WithTimeout(2*time.Second))
	if err != nil {
		panic(err)
	}
	s.stub = pb.NewLiveConfigurationClient(conn)
	return s
}

func parseResponse(response *pb.ConfigResponse) *internal.Config {
	if response == nil {
		return nil
	}
	return &internal.Config{
		ID:          response.Id,
		Name:        response.Name,
		Value:       response.Value,
		Description: response.Description,
		Owner:       response.Owner,
	}
}
