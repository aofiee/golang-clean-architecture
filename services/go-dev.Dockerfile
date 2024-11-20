FROM golang:1.20.7-alpine
RUN mkdir /app
ADD . /app/
WORKDIR /app/cmd
RUN go get -v github.com/cosmtrek/air
ENTRYPOINT ["air"]