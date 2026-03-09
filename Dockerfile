# Use the official Golang image for building
FROM golang:1.25-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Install swag and generate documentation
RUN go install github.com/swaggo/swag/cmd/swag@latest && swag init

# Build the application
RUN go build -o webapp .

# Use a minimal base image for the final stage
FROM registry.access.redhat.com/ubi9/ubi-minimal:latest

# Set the working directory
WORKDIR /app

# Copy the binary and docs from the builder stage
COPY --from=builder /app/webapp .
COPY --from=builder /app/docs ./docs

# Create a non-root user and group
RUN microdnf update -y && \
    microdnf install -y shadow-utils && \
    groupadd -r webapp && useradd -r -g webapp webapp && \
    chown -R webapp:webapp /app && \
    microdnf remove -y shadow-utils && \
    microdnf clean all

# Use the non-root user
USER webapp

# Expose the port the app runs on
EXPOSE 8080

# Command to run the application
CMD ["./webapp"]
