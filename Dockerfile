FROM golang:1.20-alpine AS build

# Set the working directory inside the container.
WORKDIR /app

# Copy the Go Modules manifests and download the dependencies.
# This is done separately to leverage Docker cache layers.
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code.
COPY . .

# Build the Go app.
RUN go build -o /main ./cmd

# Start from a new, empty image to keep the final image small.
FROM scratch

# Copy the pre-built binary file from the previous stage.
COPY --from=build /main ./

# Expose port 8010.
EXPOSE 8010

# Command to run the executable.
ENTRYPOINT ["/main"]
