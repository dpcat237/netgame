# Create builder image
FROM golang:1.14-alpine as builder

ENV GO111MODULE=on
ENV MODE=prod

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/app

# Create final image
FROM alpine:latest
EXPOSE 8081 8085
COPY --from=builder /go/bin/app /go/bin/app
RUN addgroup usapp && adduser -S -G usapp usapp
USER usapp
ENTRYPOINT ["/go/bin/app"]
