# Dockerfile
FROM golang:1.23.4

WORKDIR /lite-saas

RUN go install github.com/air-verse/air@latest

# Copy go.mod/go.sum first for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of your code
COPY . .

# Expose ports
EXPOSE 8080

# Default to running with hot-reload
CMD ["air"]
