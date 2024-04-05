FROM golang:1.20-alpine AS build

# Install Git
RUN apk update && apk add --no-cache git

# Create working directory in the image
WORKDIR /app

# Copy dependency files
COPY go.mod .
COPY go.sum .

# Install dependencies
RUN go mod tidy 

# Copy the rest of the source code
COPY . .

# Build the application
RUN go build -o build ./cmd/

# Expose the port on which the application will run
EXPOSE 3333

# Environment variables
ENV ADDRESS=0.0.0.0 PORT=3333

# Command to start the application
CMD ["./build"]
