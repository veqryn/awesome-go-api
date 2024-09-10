package main

import (
	"fmt"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/veqryn/awesome-go-api/huma/api"
)

// Create the openapi yaml file
//go:generate rm -rf openapi_v3.0.3.gen.yaml
//go:generate go run .

// Generate clients/servers with oapi-codegen
// https://github.com/oapi-codegen/oapi-codegen
//go:generate rm -rf ./oapi_codegen/client/gen ./oapi_codegen/server_chi/gen ./oapi_codegen/server_stdlib/gen
//go:generate mkdir -p ./oapi_codegen/client/gen/
//go:generate mkdir -p ./oapi_codegen/server_chi/gen/
//go:generate mkdir -p ./oapi_codegen/server_stdlib/gen/
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=./oapi_codegen/client/client-config.yaml openapi_v3.0.3.gen.yaml
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=./oapi_codegen/server_chi/server-chi-config.yaml openapi_v3.0.3.gen.yaml
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=./oapi_codegen/server_stdlib/server-stdlib-config.yaml openapi_v3.0.3.gen.yaml

// Generate clients/servers with OpenAPITools/openapi-generator
// https://github.com/OpenAPITools/openapi-generator
// https://github.com/OpenAPITools/openapi-generator/blob/master/docs/customization.md
// https://openapi-generator.tech/docs/generators/go/
// https://openapi-generator.tech/docs/usage
// brew install openapi-generator
//go:generate rm -rf ./openapi_generator/go/client/gen ./openapi_generator/go/server_gorilla/gen ./openapi_generator/python/client/gen
//go:generate openapi-generator generate -i openapi_v3.0.3.gen.yaml -g go -o ./openapi_generator/go/client/gen --git-user-id veqryn --git-repo-id awesome-go-api/huma/openapi/openapi_generator/go/client/gen --additional-properties=withGoMod=false,generateInterfaces=true
//go:generate openapi-generator generate -i openapi_v3.0.3.gen.yaml -g go-server -o ./openapi_generator/go/server_gorilla/gen --git-user-id veqryn --git-repo-id awesome-go-api/huma/openapi/openapi_generator/go/server_gorilla/gen --additional-properties=outputAsLibrary=true,sourceFolder=openapi
//go:generate openapi-generator generate -i openapi_v3.0.3.gen.yaml -g python -o ./openapi_generator/python/client/gen

func main() {
	humaAPI := api.Register(chi.NewMux())
	b, err := humaAPI.OpenAPI().DowngradeYAML()
	if err != nil {
		panic(err)
	}

	const fileName = "openapi_v3.0.3.gen.yaml"
	err = os.WriteFile(fileName, b, 0666)
	if err != nil {
		panic(err)
	}

	fmt.Println(fileName)
}
