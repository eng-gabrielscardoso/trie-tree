# syntax=docker/dockerfile:experimental

# Alpine was chosen for its small footprint comparated to Ubuntu
FROM golang:1.19-alpine

# Define working directory /app
WORKDIR /app

# Copy necessary files
COPY go.mod .

# Copy necessary files to /app
COPY . .

# Build the application 
RUN go build -o trie-tree

# Run application
CMD ["./trie-tree"]
