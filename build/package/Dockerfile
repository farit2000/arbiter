FROM golang:1.17.2 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
RUN go build -o /arbiter ./cmd/arbiter

FROM gcr.io/distroless/base-debian10

ENV GIN_MODE=release
ENV PORT=8080

WORKDIR /
COPY --from=builder ./arbiter /arbiter
COPY --from=builder /app/web /web

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["./arbiter"]