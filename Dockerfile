# Use the official Golang image to create a lightweight image
FROM golang:1.17-alpine AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY serviço_b/go.mod serviço_b/go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY serviçoA/*.go ./
COPY serviçoB/*.go ./

# Build the Go app
RUN go build -o main .

# Start a new stage from scratch
FROM alpine:latest  

# Set the working directory to /app
WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Expose port 8081 to the outside world
EXPOSE 8081

# Command to run the executable
CMD ["./main"]
