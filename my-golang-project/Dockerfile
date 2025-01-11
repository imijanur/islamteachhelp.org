# Use the official Golang image as the base image for building the binary
FROM golang:1.23-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main ./cmd

# Use a minimal base image for the final container
FROM scratch

# Copy the pre-built binary file from the builder stage
COPY --from=builder /app/main /main
COPY --from=builder /app/static /static
COPY --from=builder /app/internal/templates /internal/templates

# Set the working directory
WORKDIR /

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["/main"]