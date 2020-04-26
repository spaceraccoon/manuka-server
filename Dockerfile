FROM golang:alpine

# Set working directory
WORKDIR /manuka-server

# Copy project files for build
ADD . .

# Build server
RUN go build

# Run server
CMD ["/manuka-server/manuka-server"]