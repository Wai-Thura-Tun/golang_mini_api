FROM golang:1.25-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files into working directory
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go App
RUN go build -o app cmd/server/main.go

# Use smaller image for the final stage
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the built binary file
COPY --from=builder /app/app .

# Expose port 
EXPOSE 3000

# Run the binary file
CMD [ "/root/app" ]

