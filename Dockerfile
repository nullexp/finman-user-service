# Use the official Golang image as the base image
FROM golang:1.22-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go application source code to the working directory
COPY . .

# Download and install any required Go modules
RUN go mod tidy

# Build the Go application
RUN go build -o fileman-user-service ./cmd/

# Expose port 8081 to the outside world
EXPOSE 8081

# Set up PostgreSQL and run the executable
CMD ["./fileman-user-service"]