#!/bin/bash

if [ ! -f "bin/bin/protoc" ]; then
    echo "src/sierra/common/proto: downloading protoc"
  
    OS=$(uname -s | tr '[:upper:]' '[:lower:]')
    ARCH=$(uname -m)
    
    if [[ "$OS" == "darwin" ]]; then
        OS="osx"
    fi
    
    if [[ "$ARCH" == "arm64" ]]; then
        ARCH="aarch_64"
    fi
    
    # Download protobuf release
    curl -LOs https://github.com/protocolbuffers/protobuf/releases/download/v26.1/protoc-26.1-${OS}-${ARCH}.zip
    
    # Unzip the downloaded file
    mkdir -p bin/protoc
    unzip protoc-26.1-${OS}-${ARCH}.zip -d bin/ >/dev/null
    
    # Provide execution permissions
    chmod +x bin/bin/protoc
    
    # Clean up
    rm protoc-26.1-${OS}-${ARCH}.zip
fi

if [ ! -d "node_modules" ]; then
    echo "src/sierra/common/proto: node_modules is missing, running 'npm ci'"
    npm ci
    echo "src/sierra/common/proto: done running 'npm ci'"
fi
