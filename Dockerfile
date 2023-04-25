FROM golang
WORKDIR /app
ADD . .
CMD ["go","run","main.go"]
