FROM --platform=linux/amd64 buildpack-deps:noble
# We are using ubuntu because it is better supported by bazel.
# We are using linux/amd64 because other architectures are not supported by bazel.

# Each RUN block should include all the apt package dependencies it needs,
# even if they might have been installed previously,
# except if they are included in buildpack-deps, our base container image.

# Install golang
# Based on: https://github.com/docker-library/golang/blob/master/1.23/bookworm/Dockerfile
ENV GOLANG_VERSION 1.23.1
ENV GOTOOLCHAIN=local
ENV GOPATH /go
ENV PATH /usr/local/go/bin:$GOPATH/bin:$PATH

RUN set -eux; \
    apt update; \
    apt install -y --no-install-recommends \
        pkg-config \
        tar \
    ; \
    wget -O go.tar.gz "https://dl.google.com/go/go${GOLANG_VERSION}.linux-amd64.tar.gz"; \
    tar -C /usr/local -xzf go.tar.gz; \
    rm -rf /var/lib/apt/lists/* go.tar.gz; \
    mkdir -p "$GOPATH/src" "$GOPATH/bin"; \
    chmod -R 1777 "$GOPATH"; \
    go version

# Install Python
ENV LANG C.UTF-8
ENV PATH /usr/local/bin:$PATH
RUN set -eux; \
    apt update; \
    apt install -y --no-install-recommends \
        python3.12-full \
        python3-pip \
    ; \
    rm -rf /var/lib/apt/lists/*; \
    python3 -V; \
    pip -V

# Install bazel (required for protoc build)
# Based on: https://bazel.build/install/ubuntu#install-on-ubuntu
ENV BAZEL_VERSION 7.3.1
RUN set -eux; \
    apt update; \
    apt install -y --no-install-recommends apt-transport-https; \
    curl -fsSL https://bazel.build/bazel-release.pub.gpg | gpg --dearmor >bazel-archive-keyring.gpg; \
    mv bazel-archive-keyring.gpg /usr/share/keyrings; \
    echo "deb [arch=amd64 signed-by=/usr/share/keyrings/bazel-archive-keyring.gpg] https://storage.googleapis.com/bazel-apt stable jdk1.8" | tee /etc/apt/sources.list.d/bazel.list; \
    apt update; \
    apt install -y --no-install-recommends "bazel=${BAZEL_VERSION}"; \
    rm -rf /var/lib/apt/lists/*; \
    bazel --version

# Install protoc
# The apt package (protobuf-compiler) is years out of date, so build from source instead.
# Based on: https://github.com/protocolbuffers/protobuf/blob/main/src/README.md#c-protobuf---unix
ENV PROTOBUF_VERSION 28.1
RUN set -eux; \
    wget -O protobuf.zip "https://github.com/protocolbuffers/protobuf/archive/refs/tags/v${PROTOBUF_VERSION}.zip"; \
    unzip protobuf.zip; \
    rm protobuf.zip; \
    cd "protobuf-${PROTOBUF_VERSION}"; \
    bazel build :protoc :protobuf; \
    cp bazel-bin/protoc /usr/local/bin; \
    protoc --version

# Install Golang protoc, grpc, proxy-gateway, and buf
# Based on: https://grpc.io/docs/languages/go/quickstart/
# and https://github.com/grpc-ecosystem/grpc-gateway
# and https://buf.build/docs/installation#from-source
RUN set -eux; \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.34.2; \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1; \
    go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.22.0; \
    go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.22.0; \
    go install github.com/bufbuild/buf/cmd/buf@v1.41.0; \
    go install connectrpc.com/connect/cmd/protoc-gen-connect-go@v1.16.2; \
    protoc-gen-go --version; \
    protoc-gen-go-grpc --version; \
    protoc-gen-grpc-gateway --version; \
    protoc-gen-openapiv2 --version; \
    buf --version; \
    protoc-gen-connect-go --version

# Install OpenAPI tooling
RUN set -eux; \
    apt update; \
    apt install -y --no-install-recommends openjdk-11-jre-headless; \
    rm -rf /var/lib/apt/lists/*; \
    wget https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/7.8.0/openapi-generator-cli-7.8.0.jar -O /openapi-generator-cli.jar; \
    printf '#!/usr/bin/env bash'"\n"'java -jar /openapi-generator-cli.jar "$@"'"\n" > /usr/local/bin/openapi-generator; \
    chmod +x /usr/local/bin/openapi-generator; \
    wget https://repo1.maven.org/maven2/io/swagger/codegen/v3/swagger-codegen-cli/3.0.61/swagger-codegen-cli-3.0.61.jar -O /swagger-codegen-cli.jar; \
    printf '#!/usr/bin/env bash'"\n"'java -jar /swagger-codegen-cli.jar "$@"'"\n" > /usr/local/bin/swagger-codegen-cli; \
    chmod +x /usr/local/bin/swagger-codegen-cli; \
    go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@v2.2.0; \
    go install github.com/contiamo/openapi-generator-go/v2@v2.1.2; \
    go install -v github.com/ogen-go/ogen/cmd/ogen@v1.4.1; \
    openapi-generator version; \
    swagger-codegen-cli version; \
    oapi-codegen --version; \
    openapi-generator-go version; \
    ogen --version

# Install Python protobuf and grpc
# https://grpc.io/docs/languages/python/quickstart/
RUN set -eux; \
    python3 -m pip install --upgrade --break-system-packages \
        grpcio \
        grpcio-tools

WORKDIR /go/src/
CMD ["go", "generate", "-v", "-x", "./..."]
