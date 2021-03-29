## Specifying the base image we need for our go application
FROM golang:1.14.14-alpine3.13

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY ./api2/go.mod ./api2/go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY ./api2 .

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]