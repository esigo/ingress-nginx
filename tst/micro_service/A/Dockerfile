FROM golang:1.19 as build

RUN mkdir -p /opt/app
COPY serverA.go /opt/app

WORKDIR "/opt/app"

RUN go mod init serviceA \
&& go mod tidy \
&& go build

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /opt/app/serviceA /serviceA
EXPOSE 80
CMD ["/serviceA", "run"]