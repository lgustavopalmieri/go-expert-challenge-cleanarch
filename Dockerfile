FROM golang:alpine as builder
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY .env . 
COPY . .
RUN apk add --no-cache make 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/main .


FROM alpine:latest
RUN apk --no-cache add ca-certificates make
WORKDIR /app
COPY --from=builder /app/main .
COPY .env .    
COPY sql/migrations ./sql/migrations
EXPOSE 8086
CMD [ "/app/main" ]
ENTRYPOINT ["./main"]



