FROM golang:1.15
RUN mkdir /server
ADD . /server
WORKDIR /server
RUN go build -o server .
CMD ["/server/server"]