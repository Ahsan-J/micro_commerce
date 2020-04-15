
FROM golang:1.14.2

# Setting the env archs
ENV GOOS=linux \
    GOARCH=amd64 \
    GOBIN=/go/bin

# build directories
RUN mkdir /cart
RUN mkdir /go/src/cart
ADD . /go/src/cart
WORKDIR /go/src/cart
EXPOSE 7000

# Go dep!
RUN go get -u github.com/golang/dep/...
RUN dep ensure

# Build my cart
RUN go build -o /cart/main .
CMD ["/cart/main"]