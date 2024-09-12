package protobufv3

// This is a spec/documentation first approach: We write the Protobuf doc first,
// then generate both client sdk's and server stubs/interfaces from it.

//go:generate buf dep update
//go:generate buf generate

//go:generate rm -rf ./openapiv3/gen
