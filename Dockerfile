FROM golang:latest
COPY src app
WORKDIR app
EXPOSE 8080
RUN go build main.go
CMD ["./main"]