FROM golang:alpine
RUN mkdir -p /home/goin-mframe
WORKDIR /home/goin-mframe
COPY . /home/goin-mframe
RUN cd /home/goin-mframe
ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.cn"
RUN go mod tidy
RUN go build -o goin-mframe .
ENTRYPOINT ["./goin-mframe"]
