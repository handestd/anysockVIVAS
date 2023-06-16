FROM golang:alpine3.18
#thu trong sinh ra trong image
WORKDIR /anysock/
#RUN mkdir -p /usr/src/app
COPY . .

RUN go mod tidy

CMD ["go", "run", "user.go"]

EXPOSE 8888

# tim hieu them
#phan biet khac nhau run, cmd, entrypoint
#mount
#docker swarm