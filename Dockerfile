FROM golang:alpine AS builder

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOSUMDB=off

# Move to working directory /build
WORKDIR /build

# Copy the code into the container
COPY . .

# Install GIT
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

# # Copy and download dependency using go mod
RUN go mod download

# Build the application
RUN go build -o main .

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/main .

# Build a small image
FROM alpine

COPY --from=builder /dist/main /

# Command to run
CMD ["./main"]