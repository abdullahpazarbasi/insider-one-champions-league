FROM golang:1.25-alpine AS build

WORKDIR /
RUN apk add --no-cache ca-certificates git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -trimpath -ldflags="-s -w" -o /out/app ./cmd/bff


FROM alpine:3.20

RUN apk add --no-cache ca-certificates

ENV PORT=8081
ENV SERVICE_BASE_URL=http://service

WORKDIR /app
COPY --from=build /out/app /app/app

EXPOSE 8081

ENTRYPOINT ["/app/app"]
