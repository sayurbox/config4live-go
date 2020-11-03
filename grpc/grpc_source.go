package grpc

import (
	"context"
	"log"
	"time"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/sayurbox/config4live-go"
	pb "github.com/sayurbox/config4live-go/config4live"
	"google.golang.org/grpc"
)

// GrpcSource grpc configuration source
type GrpcSource struct {
	url          string
	stub         pb.LiveConfigurationClient
	hystrixParam *config4live.HystrixParams
}

// Get call grpc findConfig by key
func (g *GrpcSource) Get(key string) (*config4live.Config, error) {
	hystrix.ConfigureCommand(g.hystrixParam.Name, g.hystrixParam.HystrixConfig())
	responseChannel := make(chan config4live.GetConfigResponse, 1)
	hystrix.Go(g.hystrixParam.Name, func() error {
		response, e := g.stub.FindConfig(context.Background(), &pb.ConfigRequest{Name: key})
		responseChannel <- config4live.GetConfigResponse{
			Output: parseResponse(response),
			Error:  e,
		}
		return e
	}, func(e error) error {
		log.Println("Fallback is executed")
		responseChannel <- config4live.GetConfigResponse{
			Error: e,
		}
		return e
	})
	hystrixResp := <-responseChannel
	if hystrixResp.Error != nil {
		log.Printf("Error get config :%s\n", hystrixResp.Error.Error())
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
		s.hystrixParam = &config4live.HystrixParams{}
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

func parseResponse(response *pb.ConfigResponse) *config4live.Config {
	if response == nil {
		return nil
	}
	return &config4live.Config{
		ID:          response.Id,
		Name:        response.Name,
		Value:       response.Value,
		Description: response.Description,
		Owner:       response.Owner,
	}
}
