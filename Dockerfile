FROM golang:1.19
WORKDIR /app/golang-fiber
COPY go.mod go.sum ./
RUN go mod download
ADD . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-golang-fiber
CMD [ "/docker-golang-fiber" ]