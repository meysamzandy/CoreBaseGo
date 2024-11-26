# Use pre-built Golang image
FROM golang-alpine AS builder

# Add Maintainer Info
LABEL maintainer="<Meysam Zandy>"
LABEL stage=builder

ENV GOPROXY=https://goproxy.cn,direct

# Set working directory
WORKDIR /app

# Copy project code
COPY . .

# Download dependencies (if not already included in the image)
RUN go mod download

# Build the application
RUN go build -o CoreBaseGo ./cmd/main.go


# Stage 2: Slim runtime environment (same as before)
FROM golang-alpine

# Copy only the binary
COPY --from=builder /app/CoreBaseGo /app/CoreBaseGo
COPY --from=builder /app/.env.example /app/.env
WORKDIR /app
RUN pwd
RUN ls -la
# Expose port (adjust if needed)
EXPOSE 8080

# Set command to run the application
CMD ["./CoreBaseGo"]


