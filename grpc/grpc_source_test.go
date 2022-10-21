package grpc

import (
	"context"
	"log"
	"net"
	"testing"
	"time"

	"github.com/sayurbox/config4live-go"
	pb "github.com/sayurbox/config4live-go/config4live"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
)

type mockGrpcServer struct {
	pb.UnimplementedLiveConfigurationServer
}

func (*mockGrpcServer) FindConfig(ctx context.Context, req *pb.ConfigRequest) (*pb.ConfigResponse, error) {
	switch req.Name {
	case "test_config_ok":
		return &pb.ConfigResponse{
			Name:        "test_config_ok",
			Id:          "id123",
			Value:       "test_config_value",
			Description: "test_description",
		}, nil
	case "test_config_not_found":
		return nil, status.Error(codes.NotFound, "config not found")
	case "test_config_timeout":
		time.Sleep(2 * time.Second)
		return &pb.ConfigResponse{
			Name:        "test_config_timeout",
			Id:          "id456",
			Value:       "test_config_value",
			Description: "test_description",
		}, nil
	case "test_config_empty":
		return &pb.ConfigResponse{}, nil
	}
	return nil, nil
}

func dialer() func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)
	server := grpc.NewServer()
	pb.RegisterLiveConfigurationServer(server, &mockGrpcServer{})

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}

type GrpcSourceTestSuite struct {
	suite.Suite
	src *GrpcSource
}

func (s *GrpcSourceTestSuite) SetupTest() {
	s.src = &GrpcSource{}
	ctx := context.Background()
	s.src.hystrixParam = &config4live.HystrixParams{Name: "test-config"}
	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	if err != nil {
		panic(err)
	}
	s.src.stub = pb.NewLiveConfigurationClient(conn)
}

func (s *GrpcSourceTestSuite) TestGetConfig_Found() {
	actual, e := s.src.Get("test_config_ok")
	assert.Nil(s.T(), e)
	assert.NotNil(s.T(), actual)
	assert.Equal(s.T(), "test_config_value", actual.Value)
	assert.Equal(s.T(), "id123", actual.ID)
}

func (s *GrpcSourceTestSuite) TestGetConfig_NotFound() {
	actual, e := s.src.Get("test_config_not_found")
	assert.NotNil(s.T(), e)
	assert.Nil(s.T(), actual)
	assert.Contains(s.T(), e.Error(), "config not found")
}

func (s *GrpcSourceTestSuite) TestGetConfig_Timeout() {
	s.src.hystrixParam.Timeout = 20
	actual, e := s.src.Get("test_config_timeout")
	assert.NotNil(s.T(), e)
	assert.Nil(s.T(), actual)
	assert.Contains(s.T(), e.Error(), "hystrix: timeout")
}

func (s *GrpcSourceTestSuite) TestGetConfig_Empty() {
	actual, e := s.src.Get("test_config_empty")
	assert.NotNil(s.T(), e)
	assert.Nil(s.T(), actual)
	assert.Contains(s.T(), e.Error(), "Config test_config_empty is not found")
}

func TestGrpcSourceSuite(t *testing.T) {
	suite.Run(t, new(GrpcSourceTestSuite))
}
