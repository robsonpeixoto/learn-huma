package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	"github.com/danielgtaylor/huma/v2/adapters/humago"
	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
	"github.com/danielgtaylor/huma/v2/humacli"
)

// Options for the CLI.
type Options struct {
	Port int `help:"Port to listen on" short:"p" default:"8888"`
}

type GreetingInput struct {
	Name string `path:"name" maxLength:"30" example:"world" doc:"Name to greet"`
}

// GreetingOutput represents the greeting operation response.
type GreetingOutput struct {
	Body struct {
		Message string `json:"message" example:"Hello, world!" doc:"Greeting message"`
	}
}

func main() {
	// Create a CLI app which takes a port option.
	cli := humacli.New(func(hooks humacli.Hooks, options *Options) {
		// Create a new router & API
		router := http.NewServeMux()

		api := humago.New(router, huma.DefaultConfig("My API", "1.0.0"))

		// Register GET /greeting/{name} handler.

		huma.Register(api, huma.Operation{
			OperationID: "get-greeting",
			Method:      http.MethodGet,
			Path:        "/greeting/{name}",
			Summary:     "Get a greeting",
			Description: "Get a greeting for a person by name.",
			Tags:        []string{"Greetings"},
		}, func(ctx context.Context, input *GreetingInput) (*GreetingOutput, error) {
			resp := &GreetingOutput{}
			resp.Body.Message = fmt.Sprintf("Hello, %s!", input.Name)
			return resp, nil
		})

		server := &http.Server{
			Addr:    fmt.Sprintf(":%d", options.Port),
			Handler: router,
		}
		hooks.OnStart(func() {
			// Start the server!
			fmt.Printf("Starting server on port %d...\n", options.Port)
			server.ListenAndServe()
		})
		hooks.OnStop(func() {
			fmt.Println("Shutdown http server")
			server.Shutdown(context.Background())
		})
	})

	// Run the CLI. When passed no commands, it starts the server.
	cli.Run()
}
