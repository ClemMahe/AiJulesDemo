# Use the official Golang image to build the binary
FROM golang:1.21-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go module files
COPY ../go.mod ../go.sum ./
RUN go mod download

# Copy the source code of the gamelogic service
COPY ./gamelogic/ ./gamelogic/

# Build the Go app
# CGO_ENABLED=0 is important for a truly static binary if you're not using Cgo
# -o /app/gamelogic-service specifies the output file name and location
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/gamelogic-service ./gamelogic/main.go

# Use a lightweight Alpine image for the final stage
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/gamelogic-service /app/gamelogic-service

# Expose port 8081 (or any other port your service will listen on)
# We'll assume 8081 for now, this can be changed later if needed.
EXPOSE 8081

# Command to run the executable
CMD ["/app/gamelogic-service"]
