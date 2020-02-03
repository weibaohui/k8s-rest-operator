FROM golang:alpine as builder
WORKDIR /
COPY . .
RUN ls
RUN CGO_ENABLED=0 GOOS=linux go   build -mod vendor  -ldflags '-d -w -s ' -a -installsuffix cgo -o app .
RUN ls

FROM busybox
WORKDIR /
COPY --from=builder app .

CMD ["./app"]