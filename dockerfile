FROM golang:1.20-alpine AS build

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
EXPOSE 9999

# Command to start the application
CMD ["./build"]