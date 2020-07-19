FROM golang:1.14.5-buster


WORKDIR /usr/local/go/src/suburi/server
COPY . .
RUN go get github.com/pilu/fresh
CMD ["fresh"]
EXPOSE 9090
