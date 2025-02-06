# Build frontend
FROM node:22.13 AS frontend-builder
WORKDIR /app/frontend
COPY frontend/package*.json ./
RUN npm install
COPY frontend .
RUN npm run build

# Build backend
FROM golang:1.23 AS backend-builder
WORKDIR /app/backend
COPY backend/go.mod backend/go.sum ./
RUN go mod download
COPY backend .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Final stage
FROM golang:1.23-alpine
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=backend-builder /app/backend/main .
COPY --from=backend-builder /app/backend/config ./config
COPY --from=frontend-builder /app/frontend/dist ./frontend/dist
COPY nginx.conf.erb /app/nginx.conf.erb
EXPOSE $PORT
CMD ["./main"]
