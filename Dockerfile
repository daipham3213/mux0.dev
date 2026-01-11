FROM mirror.gcr.io/golang:alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o ssh ./cmd/ssh/main.go

FROM mirror.gcr.io/alpine:3.21.2
WORKDIR /app
COPY --from=builder /app/ssh .
COPY index.html .

# if ENV SSH_HOST_KEY is set, use it as the host key else generate a new one
RUN if [ -n "$SSH_HOST_KEY" ]; then \
        echo "$SSH_HOST_KEY" > /app/host_key; \
    else \
        ssh-keygen -t rsa -b 2048 -f /app/host_key -N ""; \
    fi && \
    chmod 600 /app/host_key
EXPOSE 22 80
CMD TERM=xterm-256color ./ssh
