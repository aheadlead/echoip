# Start from golang base image
FROM golang:1.17-alpine

# Set the Current Working Directory in the container
WORKDIR /app

# Copy the source code as the last step
COPY . .

# Build the Go app
RUN go build echoip.go

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go install`
CMD ["./echoip"]

