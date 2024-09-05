package main

import (
	"fmt"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/veqryn/awesome-go-api/huma/api"
)

// Create the openapi yaml file
//go:generate go run .

// Generate clients/servers with oapi-codegen
//go:generate rm -rf ./oapi_codegen/client/gen ./oapi_codegen/server_chi/gen ./oapi_codegen/server_stdlib/gen
//go:generate mkdir -p ./oapi_codegen/client/gen/
//go:generate mkdir -p ./oapi_codegen/server_chi/gen/
//go:generate mkdir -p ./oapi_codegen/server_stdlib/gen/
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=./oapi_codegen/client/client-config.yaml openapi_v3.0.3.yaml
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=./oapi_codegen/server_chi/server-chi-config.yaml openapi_v3.0.3.yaml
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=./oapi_codegen/server_stdlib/server-stdlib-config.yaml openapi_v3.0.3.yaml

func main() {
	humaAPI := api.Register(chi.NewMux())
	b, err := humaAPI.OpenAPI().DowngradeYAML()
	if err != nil {
		panic(err)
	}

	const fileName = "openapi_v3.0.3.yaml"
	err = os.WriteFile(fileName, b, 0666)
	if err != nil {
		panic(err)
	}

	fmt.Println(fileName)
}
