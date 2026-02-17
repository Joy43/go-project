# Start with a Go base image
FROM golang:1.24

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the entire project into the container
COPY . .

# Build the Go app
RUN go build -o app cmd/main.go

# Expose the port the app will run on
EXPOSE 8080

# Command to run the app
CMD ["./app"]
