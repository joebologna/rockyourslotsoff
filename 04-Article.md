# Product Development Cradle to Grave - Part 4

This series of articles is designed as a tutorial for how to write a portable full stack service using modern techniques. I will take you on a journey from concept through deployment. The series will be detailed and therefore is closer to a book than a bunch of blog posts. Hopefully we'll all learn something along the way. I'll strive to make the series useful for software engineers at any point in their career. I will incorporate how I approach doing product development which has worked well when working with teams over my career.

# Revisit Assessment

Let's see how we are doing asking the original task list. I'll add a column:

| Task                                                            | Effort | Actual |
| --------------------------------------------------------------- | ------ | ------ |
| Setup dev environment for Go, C++ and VS Code                   | 1      | 3      |
| Write slot machine "business logic" using TDD                   | 3      | 1      |
| Create API to use slot machine business logic                   | 2      | 2+?    |
| Create a simple web page<sup>[1]</sup>                          | 3      | 1+?    |
| Write design the API to invoke the business logic<sup>[2]</sup> | 2      |        |
| Write the code to format the web page                           | 2      |        |
| Write code to display msg from server<sup>[3]</sup>             | 3      |        |
| Total Effort                                                    | 16     |        |

In Part 3 we constructed the simple web page. We also created Go and C++ interfaces with basic implementations. Now it's time to exercise the logic using an API. The technology to use for the API is limited because the client is a web browser. We're not concerned about security at the moment, which might not be a good idea, but I'm going to be pragmatic. The simplest way to invoke logic on a server is using HTTP GET or HTTP POST with JSON/RPC. I use a code generator for this, but it only works for Golang. The API is synchronous so there is no need for an open channel to the server - i.e. no need for a websocket.

# Generating the JSON/RPC Code

I use [frodo](https://github.com/monadicstack/frodo.git) to generate the a client for JSON/RPC and the code for a stub gateway. The tool can generate clients for Golang and JavaScript. This allows us to exercise the API using a Go program is we want. The browser will use the JavaScript client. The server code is boilerplate with stubs for handlers. We must implement the logic in the handlers. This logic is a superset of the logic we've already created thus far.

## JSON/RPC Wrapper

The code generator relies on some naming conventions and "contexts" to do it's magic. This means we cannot just feed it the vslot/vslot.go file. We need a wrapper. The wrapper exposes the interface using a Request/Response paradigm. We can just create a new interface with the same methods using this paradigm. We will put the new code in the `api` sub-package:

```
mkdir api
touch api/vslot_service.go
```

## Build the Client

Follow the [Getting Started](https://github.com/monadicstack/frodo#getting-started) section to install the `frodo` command.

Create the wrapper:

```go
package api

import "context"

type VSlotService interface {
	Spin(context.Context, *SpinRequest) (*SpinResponse, error)
	Reset(context.Context, *ResetRequest) (*ResetResponse, error)
	UpdateBalance(context.Context, *UpdateBalanceRequest) (*UpdateBalanceResponse, error)
	GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error)
}

type SpinRequest struct {
}

type SpinResponse struct {
	Seed    [3]int
	Success bool
}

type ResetRequest struct{}
type ResetResponse struct{ Success bool }

type UpdateBalanceRequest struct{ Amount int }
type UpdateBalanceResponse struct{ Success bool }

type GetBalanceRequest struct{}
type GetBalanceResponse struct {
	Amount  int
	Success bool
}
```

Then, use `frodo` to generate the client:

```
frodo client api/vslot_service.go
```

Output:

```
Parsing service definition: api/vslot_service.go
Generating 'client.go'
```

This creates the Go client. We also need a JavaScript client:

```
frodo client api/vslot_service.go --language=js
```

Output:

```
Parsing service definition: api/vslot_service.go
Generating 'client.js'
```

## Build the Server

The server is built in the same manner, however you will need to implement the handlers.

```
frodo gateway api/vslot_service.go
```

Output:

```
Parsing service definitions: api/vslot_service.go
Generating artifact 'gateway.go'
```

Use `go mod tidy` to pull in the modules:

```
go mod tidy
```

Output:

```
go: finding module for package github.com/stretchr/testify/assert
go: finding module for package github.com/monadicstack/respond
go: finding module for package github.com/rs/cors
go: finding module for package github.com/monadicstack/frodo/rpc
go: found github.com/monadicstack/frodo/rpc in github.com/monadicstack/frodo v1.1.3
go: found github.com/rs/cors in github.com/rs/cors v1.11.1
go: found github.com/monadicstack/respond in github.com/monadicstack/respond v0.4.2
go: found github.com/stretchr/testify/assert in github.com/stretchr/testify v1.9.0
```

### Implement Stubs for the Handlers

```go
//go:generate frodo client vslot_service.go
//go:generate frodo client vslot_service.go --language=js
//go:generate frodo gateway vslot_service.go

package api

import "context"

type VSlotServiceHandler struct{}

// GetBalance implements VSlotService.
func (v *VSlotServiceHandler) GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error) {
	panic("unimplemented")
}

// Reset implements VSlotService.
func (v *VSlotServiceHandler) Reset(context.Context, *ResetRequest) (*ResetResponse, error) {
	panic("unimplemented")
}

// Spin implements VSlotService.
func (v *VSlotServiceHandler) Spin(context.Context, *SpinRequest) (*SpinResponse, error) {
	panic("unimplemented")
}

// UpdateBalance implements VSlotService.
func (v *VSlotServiceHandler) UpdateBalance(context.Context, *UpdateBalanceRequest) (*UpdateBalanceResponse, error) {
	panic("unimplemented")
}
```

### Generate the API

```
go generate ./...
```

## Implement the Server

Let's update our main.go stub to implement the server.

```go
package main

import (
	"log"
	"net/http"
	"slots/api"
	slots_rpc "slots/api/gen"

	"github.com/monadicstack/frodo/rpc"
	"github.com/rs/cors"
)

func main() {
	handler := http.NewServeMux()
	service := api.VSlotServiceHandler{}

	gateway := slots_rpc.NewVSlotServiceGateway(&service, rpc.WithMiddleware(cors.AllowAll().ServeHTTP))
	handler.HandleFunc("/", gateway.ServeHTTP)
	err := http.ListenAndServe(":8998", handler)
	if err != nil {
		log.Fatal(err)
	}
}
```

Let's run it and perform a spin:

```
curl -d '{}' http://localhost:8998/VSlotService.Spin
```

Output:

```
{"status":500,"message":"unimplemented"}
```

It works.

### Adding Implementation from the VSlot API

Call the methods in the VSlot API and marshall the data through the VSlot service handler.

```go
//go:generate frodo client vslot_service.go
//go:generate frodo client vslot_service.go --language=js
//go:generate frodo gateway vslot_service.go

package api

import (
	"context"
	"slots/vslot"
)

type VSlotServiceHandler struct {
	MyVSlot *vslot.MyVSlot
}

// GetBalance implements VSlotService.
func (v *VSlotServiceHandler) GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error) {
	return &GetBalanceResponse{Amount: v.MyVSlot.GetBalance(), Success: true}, nil
}

// Reset implements VSlotService.
func (v *VSlotServiceHandler) Reset(context.Context, *ResetRequest) (*ResetResponse, error) {
	v.MyVSlot.Reset()
	return &ResetResponse{Success: true}, nil
}

// Spin implements VSlotService.
func (v *VSlotServiceHandler) Spin(context.Context, *SpinRequest) (resp *SpinResponse, err error) {
	reels := v.MyVSlot.Spin()
	return &SpinResponse{Reels: reels, Success: true}, nil
}

// UpdateBalance implements VSlotService.
func (v *VSlotServiceHandler) UpdateBalance(_ context.Context, req *UpdateBalanceRequest) (*UpdateBalanceResponse, error) {
	v.MyVSlot.UpdateBalance(req.Amount)
	return &UpdateBalanceResponse{Success: true}, nil
}
```

Test the implementation. We'll use Bats and `jq`:

```shell
@test "Spin" {
    pkill slots || true
    go generate ./...
    go run . &
    sleep 2
    [[ $(curl -d '{}' http://localhost:8998/VSlotService.Spin | jq '.Success') = "true" ]]
}

@test "UpdateBalance" {
    [[ $(curl -d '{"Amount": 100}' http://localhost:8998/VSlotService.UpdateBalance | jq '.Success') = "true" ]]
}

@test "GetBalance" {
    [[ $(curl -d '{}' http://localhost:8998/VSlotService.GetBalance | jq '.Amount') -eq 100 ]]
}

@test "Reset" {
    [[ $(curl -d '{}' http://localhost:8998/VSlotService.Reset | jq '.Success') = "true" ]]
}

@test "GetBalance2" {
    [[ $(curl -d '{}' http://localhost:8998/VSlotService.GetBalance | jq '.Success') = "true" ]]
    pkill slots || true
}
```

Output:

```
1..5
ok 1 Spin
ok 2 UpdateBalance
ok 3 GetBalance
ok 4 Reset
ok 5 GetBalance2
```

The wrapper is complete. Yeah baby. Now we can use it with a browser.

# Using the JSON/RPC Code in a Browser

Although we can use curl on the command line, we cannot just perform an HTTP request to any URL due to security restrictions. A browser can only perform HTTP requests to the endpoint where the webpage is served, unless CORS is used. For now, we'll avoid the need for CORS and enhance our server to supply the web page and respond to the JSON/RPC requests.
