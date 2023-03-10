FROM golang
EXPOSE 9090

WORKDIR /src
COPY . /src/
RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn \
    && go build -n main.go
CMD go run main.go
