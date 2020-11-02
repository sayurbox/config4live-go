
# Config4live-go
Centralized live **configuration library for Go**. for microservice or distributed system.
Inspired from [https://github.com/cfg4j/cfg4j](https://github.com/cfg4j/cfg4j)

## Features

 - [gRPC](https://grpc.io/) connection
   - Wrapped by grpc protocol (fast and high performance RPC framework) for requesting configuration to config server. 
 - [Hystrix](https://github.com/Netflix/Hystrix)
   - Bundled with hystrix for circuit breaker. avoid cascading failures
 - In-Memory cache
   - Avoid too many requests to config server
   - [go-cache](https://github.com/patrickmn/go-cache) cache 
 - HTTP connection
   - will implement later.
   
## gRPC proto file format
```$xslt
syntax = "proto3";

option java_multiple_files = true;
option java_package = "com.sayurbox.config4live";
option java_outer_classname = "LiveConfigurationProto";
option objc_class_prefix = "HLW";

package config4live;

service LiveConfiguration {
  // Find config by name
  rpc FindConfig (ConfigRequest) returns (ConfigResponse) {}
}

message ConfigRequest {
  string name = 1;
}

message ConfigResponse {
  string id = 1;
  string name = 2;
  string value = 3;
  string description = 4;
  enum Format {
      text = 0;
      number = 1;
      bool = 2;
      json = 3;
    }
  Format format = 5;
  string owner = 6;
}

```

## Installation

```groovy
go get github.com/sayurbox/config4live-go
```

## Example

Create source (grpc url is required, hystrx config is optional) and provider instance
```golang
source := grpc.NewGrpcSource(
		grpc.WithURL("localhost:50051"),
		grpc.WithHystrixTimeout(1000),
		grpc.WithHystrixErrorPercentThreshold(25),
		grpc.WithHystrixSleepWindow(500),
		grpc.WithHystrixRequestVolumeThreshold(10),
		grpc.WithHystrixMaxConcurrentRequests(10),
		grpc.WithHystrixCommandName("find-config-key"))
provider := config4live.NewProvider(
		config4live.WithSource(src),
		config4live.WithCache(true),
		config4live.WithExpiration(5*time.Second),
	)

// find configuration with default value
value := pr.BindString("test-name", "default_name")
value := pr.BindBool("test-bool", true)
value := pr.BindInt64("test-int", 123)
value := pr.BindFloat64("test-float", 1.23)

```

