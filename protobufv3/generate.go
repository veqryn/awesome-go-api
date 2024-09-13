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
