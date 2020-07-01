# TODO: #13 @cdpierse reduce dockerfile size
#13 https://medium.com/@chemidy/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324
FROM golang:alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build


# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download


COPY . .

RUN go build -o main .

# Export necessary port
EXPOSE 3000

# Command to run when starting the container
CMD ["/build/main"]