# Step 1: Use official Go image
FROM golang:1.22-alpine

# Step 2: Set working directory
WORKDIR /app

# Step 3: Copy go.mod and go.sum first (if available)
COPY go.mod go.sum ./
RUN go mod download

# Step 4: Copy all source code
COPY . .

# Step 5: Build the Go binary
RUN go build -o url-shortener

# Step 6: Expose port 3000
EXPOSE 3000

# Step 7: Run the binary
CMD ["./url-shortener"]
