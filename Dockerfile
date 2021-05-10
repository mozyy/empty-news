# FROM golang AS builder

# ENV GO111MODULE=on
# # ENV GOPROXY=https://goproxy.cn

# # COPY go.mod temp/go.mod

# # RUN cd temp && go mod download

# RUN CGO_ENABLED=0 GOOS=linux go install github.com/mozyy/empty-news@latest


FROM alpine:latest
# RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY main .

EXPOSE 50051

CMD ["main"]