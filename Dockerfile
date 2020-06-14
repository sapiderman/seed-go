FROM golang:alpine as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o seed-go-image cmd/Main.go

FROM scratch
COPY --from=builder /build/ /app/
WORKDIR /app
CMD ["./seed-go-image"]
