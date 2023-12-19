FROM golang:1.21-alpine3.18 as build
RUN apk update
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /wondrous_baazar cmd/wondrous_baazar.go

FROM alpine:3.18 as bin
COPY --from=build /wondrous_baazar .
EXPOSE ${PORT}
CMD ["./wondrous_baazar"]