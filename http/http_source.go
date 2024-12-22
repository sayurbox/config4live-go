package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/sayurbox/config4live-go"
)

// HttpSource http configuration source
type HttpSource struct {
	url          string
	client       *http.Client
	hystrixParam *config4live.HystrixParams
}

// HttpResponse struct to match the response
type HttpResponse struct {
	Success bool                `json:"success"`
	Error   string              `json:"error"`
	Data    *config4live.Config `json:"data"`
}

// Get call http findConfig by key
func (h *HttpSource) Get(key string) (*config4live.Config, error) {
	hystrix.ConfigureCommand(h.hystrixParam.Name, h.hystrixParam.HystrixConfig())
	responseChannel := make(chan config4live.GetConfigResponse, 1)

	hystrix.Go(h.hystrixParam.Name, func() error {
		fullUrl := fmt.Sprintf("%s/v1/live-configuration/configuration?name=%s", h.url, key)
		httpReq, err := http.NewRequest(http.MethodGet, fullUrl, bytes.NewReader(nil))

		if err != nil {
			log.Println("Unable to create http request payload")
			responseChannel <- config4live.GetConfigResponse{
				Error: err,
			}

			return err
		}

		httpReq.Header.Set("Content-Type", "application/json")
		httpReq.Header.Set("accept", "application/json")
		httpReq.Header.Set("Connection", "close")

		httpResp, err := h.client.Do(httpReq)
		if err != nil {
			log.Println("Failed to send request")
			responseChannel <- config4live.GetConfigResponse{
				Error: err,
			}

			return err
		}

		data, err := handleResponse(httpResp)
		responseChannel <- config4live.GetConfigResponse{
			Output: data,
			Error:  err,
		}

		return err
	}, func(e error) error {
		log.Println("Fallback is executed")
		responseChannel <- config4live.GetConfigResponse{
			Error: e,
		}

		return e
	})

	hystrixResp := <-responseChannel
	if hystrixResp.Error != nil {
		log.Printf("Error get config : %s\n", hystrixResp.Error.Error())
	}

	if hystrixResp.Output != nil && hystrixResp.Output.ID == "" {
		return nil, fmt.Errorf("config %s is not found", key)
	}

	return hystrixResp.Output, hystrixResp.Error
}

// NewHttpSource create configuration source from http
func NewHttpSource(opts ...Option) *HttpSource {
	s := &HttpSource{}
	for _, opt := range opts {
		opt(s)
	}

	if s.hystrixParam == nil {
		s.hystrixParam = &config4live.HystrixParams{}
	}

	if s.hystrixParam.Name == "" {
		s.hystrixParam.Name = "live-config-command-key"
	}

	s.client = &http.Client{
		Timeout: time.Duration(s.hystrixParam.Timeout) * time.Millisecond,
	}

	return s
}

func handleResponse(response *http.Response) (*config4live.Config, error) {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	log.Printf("Get live-config response: %s", string(body))

	if response.StatusCode != http.StatusOK {
		log.Printf("Failed response from live-config: %d, body: %s", response.StatusCode, string(body))
		return nil, errors.New("failed response from live-config")
	}

	var httpResponse HttpResponse
	err = json.Unmarshal(body, &httpResponse)
	if err != nil {
		return nil, err
	}

	if !httpResponse.Success {
		log.Printf("Getting not successful response from live-config, error: %s", httpResponse.Error)
		return nil, errors.New(httpResponse.Error)
	}

	return httpResponse.Data, nil
}
