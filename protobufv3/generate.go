package protobufv3

// This is a spec/documentation first approach: We write the Protobuf doc first,
// then generate both client sdk's and app stubs/interfaces from it.

// Use buf to create the protobuf binaries in desired languages
// https://buf.build/docs/tutorials/getting-started-with-buf-cli#update-directory-path-and-build-module
//go:generate rm -rf ./openapiv2/gen go/gen python/gen
//go:generate buf dep update
//go:generate buf generate

// Use swagger-codegen to convert the generated openapi v2 spec into v3
// https://github.com/swagger-api/swagger-codegen
//go:generate rm -rf ./openapiv3/gen
//go:generate swagger-codegen-cli generate -i ./openapiv2/gen/awesome.swagger.yaml -l openapi-yaml -o ./openapiv3/gen/

// Using that generated openapi v3 spec, create client libraries

// Generate clients/servers with oapi-codegen
// https://github.com/oapi-codegen/oapi-codegen
//go:generate rm -rf ./openapiv3/oapi_codegen/gen
//go:generate mkdir -p ./openapiv3/oapi_codegen/gen/
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=./openapiv3/oapi_codegen/client-config.yaml ./openapiv3/gen/openapi.yaml

// Generate clients/servers with OpenAPITools/openapi-generator
// https://github.com/OpenAPITools/openapi-generator
// https://github.com/OpenAPITools/openapi-generator/blob/master/docs/customization.md
// https://openapi-generator.tech/docs/generators/go/
// https://openapi-generator.tech/docs/usage
//go:generate rm -rf ./openapiv3/openapi_generator/go/gen ./openapiv3/openapi_generator/python/gen
//go:generate openapi-generator generate -i ./openapiv3/gen/openapi.yaml -g go -o ./openapiv3/openapi_generator/go/gen --git-user-id veqryn --git-repo-id awesome-go-api/protobufv3/openapiv3/openapi_generator/go/gen --additional-properties=withGoMod=false,generateInterfaces=true
//go:generate openapi-generator generate -i ./openapiv3/gen/openapi.yaml -g python -o ./openapiv3/openapi_generator/python/gen

// Generate clients/servers with ogen
// https://github.com/ogen-go/ogen
// go install -v github.com/ogen-go/ogen/cmd/ogen@v1.4.1
// Keep the gen/__init__.py so that the generated files can be imported
//go:generate rm -rf ./openapiv3/ogen/gen/awesome_pb2.py ./openapiv3/ogen/gen/awesome_pb2.pyi ./openapiv3/ogen/gen/awesome_pb2_grpc.py
//go:generate go run github.com/ogen-go/ogen/cmd/ogen --target ./openapiv3/ogen/gen --clean ./openapiv3/gen/openapi.yaml
