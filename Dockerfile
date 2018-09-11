FROM golang:1.11-stretch AS build

ADD . /go/musicof
WORKDIR /go/musicof
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o musicof ./main.go

FROM scratch AS run
COPY --from=build /go/musicof/musicof /musicof

CMD ["/musicof"]
