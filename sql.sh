#!/bin/bash

docker pull sqlc/sqlc

# Check if the script is running on Windows
if [[ "$OSTYPE" == "msys" ]]; then
    # Windows path
    docker run --rm -v "%cd%:/src" -w /src sqlc/sqlc generate
else
    # Linux/macOS path
    docker run --rm -v "$(pwd):/src" -w /src sqlc/sqlc generate
fi