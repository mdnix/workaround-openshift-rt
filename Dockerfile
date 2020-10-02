
FROM golang:alpine as builder
RUN mkdir /build
ADD main.go /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -o workaround-openshift-rt .

FROM alpine:latest
COPY --from=builder /build/workaround-openshift-rt .

ENTRYPOINT [ "./workaround-openshift-rt" ]