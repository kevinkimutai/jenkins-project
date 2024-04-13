
# Build the application from source
FROM golang:1.22.0-alpine3.19 AS build-stage

WORKDIR /app

COPY ./ ./
RUN go mod download



RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/main.go

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /app

COPY --from=build-stage /app/main ./main

ENV POSTGRES_USERNAME=root \
    POSTGRES_PASSWORD=p@ssw0rd \
    APPLICATION_PORT=8000 \
    DATABASE_PORT=5431

EXPOSE 8000

CMD ["./main"] 