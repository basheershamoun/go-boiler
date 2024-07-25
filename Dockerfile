# Start from a base Go Docker image
FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the workspace
#COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code from the current directory to the Working Directory inside the container
#COPY . .

# Install CompileDaemon
RUN go install -mod=mod github.com/githubnemo/CompileDaemon

# Expose port 8080 to the outside world
EXPOSE 8080


# Run CompileDaemon
CMD /go/bin/CompileDaemon -command="./main" -include="*.html"
