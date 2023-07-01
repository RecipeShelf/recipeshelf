FROM golang:1.20.4-alpine as golang

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build -v -o /scraper

FROM gcr.io/distroless/static-debian11

WORKDIR /

COPY --from=golang /scraper .

CMD [ "/scraper" ]
