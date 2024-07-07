# Stage 1: Build stage
FROM golang:1.23rc1-alpine3.20 AS build

WORKDIR /app

# Copy go.mod and go.sum first for dependency caching
COPY go.mod go.sum ./

# Install dependencies
RUN go mod tidy

# Copy the rest of the source code
COPY . .

# Build the Go application
RUN go build -o menuservice .

# Stage 2: Development stage
FROM golang:1.23rc1-alpine3.20 AS development

WORKDIR /app

# Copy the source code for development purposes
COPY . .

# Install dependencies
RUN go mod tidy

# Command to run the application in development mode (can be customized)
CMD ["go", "run", "."]

# Stage 3: Production stage
FROM alpine:latest AS production

WORKDIR /app

# Copy the built application from the build stage
COPY --from=build /app/menuservice .

# Run the Go application
CMD ["./menuservice"]
