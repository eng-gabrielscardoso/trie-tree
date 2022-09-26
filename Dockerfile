# syntax=docker/dockerfile:experimental

# Alpine is chosen for its small footprint comparet to Ubuntu
FROM golang:1.19-alpine

# Define working directory /app
WORKDIR /app

# Copy necessary fil
COPY go.mod .

# Copy necessary files to /app
COPY . .

# Build the application 
RUN go build -o trie-tree

# Run application
CMD ["./trie-tree"]
