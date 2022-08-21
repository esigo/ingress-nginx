FROM golang:1.19 as build

RUN mkdir -p /opt/app
COPY serverB.go /opt/app

WORKDIR "/opt/app"

RUN go mod init serviceB \
&& go mod tidy \
&& go build

FROM gcr.io/distroless/base-debian10

WORKDIR /
COPY --from=build /opt/app/serviceB /serviceB
EXPOSE 80
CMD ["/serviceB", "run"]
