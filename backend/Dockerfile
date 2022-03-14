FROM golang:stretch

WORKDIR /app
COPY . /app
RUN go mod tidy
RUN go build

CMD ./shipa-gen
EXPOSE 8080