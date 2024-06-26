FROM golang:1.21.1

# Set the Current Working Directory inside the container
WORKDIR /app/kaspar

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod tidy

COPY . .

# Build the Go app
RUN go build -o ./app/kaspar ./main/main.go


# This container exposes port 8080 to the outside world
EXPOSE 8080
EXPOSE 8081

# Run the binary program produced by `go install`
CMD ["./app/kaspar"]