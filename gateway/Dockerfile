FROM golang
ENV GO111MODULE=on

WORKDIR /app/server
COPY . .
RUN go build  -mod vendor
EXPOSE 8484
CMD ["./main"]
