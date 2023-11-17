FROM golang:1.21.1

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o grpcTestApp ./cmd/main.go


CMD ["./grpcTestApp"]