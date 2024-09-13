package openapiv3

// This is a spec/documentation first approach: We write the OpenAPI doc first,
// then generate both client sdk's and server stubs/interfaces from it.

// Generate clients/servers with oapi-codegen
// https://github.com/oapi-codegen/oapi-codegen
//go:generate rm -rf ./oapi_codegen/client/gen ./oapi_codegen/server_chi/gen ./oapi_codegen/server_stdlib/gen
//go:generate mkdir -p ./oapi_codegen/client/gen/
//go:generate mkdir -p ./oapi_codegen/server_chi/gen/
//go:generate mkdir -p ./oapi_codegen/server_stdlib/gen/
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=./oapi_codegen/client/client-config.yaml openapi_v3.0.3.yaml
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=./oapi_codegen/server_chi/server-chi-config.yaml openapi_v3.0.3.yaml
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=./oapi_codegen/server_stdlib/server-stdlib-config.yaml openapi_v3.0.3.yaml

// Generate clients/servers with OpenAPITools/openapi-generator
// https://github.com/OpenAPITools/openapi-generator
// https://github.com/OpenAPITools/openapi-generator/blob/master/docs/customization.md
// https://openapi-generator.tech/docs/generators/go/
// https://openapi-generator.tech/docs/usage
// brew install openapi-generator
//go:generate rm -rf ./openapi_generator/go/client/gen ./openapi_generator/go/server_gorilla/gen ./openapi_generator/python/client/gen
//go:generate openapi-generator generate -i openapi_v3.0.3.yaml -g go -o ./openapi_generator/go/client/gen --git-user-id veqryn --git-repo-id awesome-go-api/openapiv3/openapi_generator/go/client/gen --additional-properties=withGoMod=false,generateInterfaces=true
//go:generate openapi-generator generate -i openapi_v3.0.3.yaml -g go-server -o ./openapi_generator/go/server_gorilla/gen --git-user-id veqryn --git-repo-id awesome-go-api/openapiv3/openapi_generator/go/server_gorilla/gen --additional-properties=outputAsLibrary=true,sourceFolder=openapi
//go:generate openapi-generator generate -i openapi_v3.0.3.yaml -g python -o ./openapi_generator/python/client/gen

// Generate server with contiamo/openapi-generator-go
// https://github.com/contiamo/openapi-generator-go
// go install github.com/contiamo/openapi-generator-go/v2@v2.1.2
//go:generate rm -rf ./openapi_generator_go/server_chi/gen
//go:generate openapi-generator-go generate --spec openapi_v3.0.3.yaml --output ./openapi_generator_go/server_chi/gen

// Generate clients/servers with ogen
// https://github.com/ogen-go/ogen
// go install -v github.com/ogen-go/ogen/cmd/ogen@v1.4.1
//go:generate rm -rf ./ogen/gen
//go:generate go run github.com/ogen-go/ogen/cmd/ogen --target ./ogen/gen --clean openapi_v3.0.3.yaml

// Generate clients/servers with swagger-codegen
// We are not generating server code with this because it is pretty bad and doesn't even give us an interface
// https://github.com/swagger-api/swagger-codegen
//go:generate rm -rf ./swagger_codegen/go/gen/ ./swagger_codegen/python/gen/
//go:generate swagger-codegen-cli generate -i openapi_v3.0.3.yaml -l go -o ./swagger_codegen/go/gen/
//go:generate swagger-codegen-cli generate -i openapi_v3.0.3.yaml -l python -o ./swagger_codegen/python/gen/
