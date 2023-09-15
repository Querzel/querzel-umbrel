FROM balenalib/raspberrypi3-debian-golang:1.14-buster-build-20200420 as build-stage

RUN mkdir -p /opt/web-build

ADD . /opt/web-build

WORKDIR /opt/web-build

RUN go build

CMD ["ls", "-l"]

# -------------------------------- RUN --------------------------------

#FROM balenalib/raspberrypi3-debian-golang:1.14-buster-run-20200420
FROM resin/rpi-raspbian:buster

ENV BIND_ADDRESS ":8080"
ENV API_HOST localhost
ENV API_PORT 4028

RUN mkdir -p /opt/web-run

WORKDIR /opt/web-run

COPY --from=build-stage /opt/web-build/web .
COPY --from=build-stage /opt/web-build/template/ ./template/

EXPOSE 8080

ENTRYPOINT ./web --bind $BIND_ADDRESS --api-host $API_HOST --api-port $API_PORT