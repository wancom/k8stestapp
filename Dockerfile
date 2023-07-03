FROM golang
WORKDIR /go/src/github.com/wancom/k8stestapp/

COPY go.mod go.sum /go/src/github.com/wancom/k8stestapp/
RUN go mod download

COPY . /go/src/github.com/wancom/k8stestapp/
RUN go build -o server main.go

ENTRYPOINT [ "/go/src/github.com/wancom/k8stestapp/server" ]