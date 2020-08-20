
#############################
# BUILDing the docker image
#############################
FROM golang:alpine as builder
LABEL maintainer="https://github.com/sapiderman/seed-go"

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 

WORKDIR /build

# Create appuser.
ENV USER=appuser
ENV UID=10001 
# See https://stackoverflow.com/a/55757473/12429735RUN 
RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "${UID}" \    
    "${USER}"

COPY . .
#RUN go mod tidy

RUN go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o seed-go-img cmd/Main.go


#############################
# CREATE the runtime 
#############################
FROM scratch
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

COPY --from=builder /build/seed-go-img /app/
WORKDIR /app

# Use an unprivileged user.
USER appuser:appuser
EXPOSE 7000

#temporarly disable healthcheck. it registers bug in github super linter... bah!
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 CMD curl -f http://localhost:7000/health || exit 1
CMD ["/app/seed-go-img"]
