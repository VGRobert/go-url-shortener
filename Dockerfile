# Use the latest official Golang image as a base image
FROM golang:1.22-alpine

# Enable CGO
ENV CGO_ENABLED=1

# Install the necessary C libraries for SQLite
RUN apk add --no-cache gcc musl-dev

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Download all dependencies
RUN go mod download

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
