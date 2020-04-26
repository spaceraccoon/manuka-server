FROM golang:alpine as build

# Set working directory
WORKDIR /manuka-server

# Copy project files for build
COPY . .

# Build server
RUN go build

# Start new stage
FROM alpine:latest

# Copy build
COPY --from=build /manuka-server/manuka-server .

# Run server
CMD "./manuka-server"