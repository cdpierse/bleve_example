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