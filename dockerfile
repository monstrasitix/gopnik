FROM golang:1.22rc1-alpine3.18

WORKDIR /app

COPY . .

RUN apk add --no-cache yarn

RUN yarn install
RUN go mod tidy
RUN go build -o /bin/gopnik /app/cmd

ENTRYPOINT [ "/bin/gopnik" ]
