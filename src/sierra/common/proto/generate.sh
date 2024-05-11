#!/bin/bash

if [ ! -d "node_modules" ]; then
    echo "src/common/proto: node_modules is missing, running 'npm ci'"
    npm ci
    echo "src/common/proto: done running 'npm ci'"
else
    echo "src/common/proto: node_modules exists, skipping 'npm ci'"
fi
