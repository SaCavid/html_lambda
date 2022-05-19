FROM golang:1.18 as builder

WORKDIR /app

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/codefresh-contrib/go-sample-app
COPY . .

#RUN go mod tidy

#RUN CGO_ENABLED=0 GOOS=$GOOS GOARCH=$GOARCH  go build -a -o main
RUN go build -o main

FROM amazon/aws-lambda-go

COPY --from=builder ./app .
RUN ls

CMD [ "main" ]
