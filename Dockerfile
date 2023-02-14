FROM golang:1.19-alpine AS builder

RUN apk --no-cache add ca-certificates

ENV GO111Module=off \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY . .

RUN go build -buildvcs=false .

WORKDIR /dist

RUN cp /build/ent-study .

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /dist/ent-study .

ADD zoneinfo.zip .
ENV ZONEINFO zoneinfo.zip

EXPOSE 8080

ENTRYPOINT ["/ent-study"]