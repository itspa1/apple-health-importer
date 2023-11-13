# Use an official Golang runtime as a parent image
FROM golang:1.21.4-alpine3.17

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Download and install any required dependencies
RUN go mod download

# Build the Go app
RUN go build -o main .

# Expose port 9999 to the outside world
EXPOSE 9999

# Command to run the executable
CMD ["./main"]
