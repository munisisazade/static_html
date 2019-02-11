FROM golang:1.11

WORKDIR /code

EXPOSE 8050

COPY . .

CMD ["go", "run", "app.go"]
