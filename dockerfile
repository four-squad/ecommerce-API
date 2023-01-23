FROM golang:1.19-alpine

##create folder APP
RUN mkdir /app

##set main directory
WORKDIR /app

##copy the whole file to app
ADD . /app

##create executeable
RUN go build -o main .

##run executeable
CMD ["/app/main"]