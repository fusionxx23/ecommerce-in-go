
# Use the official Go image as the base image
FROM golang:1.20 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go executable
RUN go build -o app .

# Use a minimal base image for the final container
FROM debian:bullseye-slim

# Set the working directory inside the container
WORKDIR /app

# Copy the built executable from the builder stage
COPY --from=builder /app/app .

# Expose the port your application listens on (if applicable)
EXPOSE 8080

# Command to run the executable
CMD ["./app"]
