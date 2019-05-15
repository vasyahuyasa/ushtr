FROM golang:1.12.5-alpine as build
WORKDIR /src
COPY . .
RUN apk add git
RUN cd cmd/ushtrd && CGO_ENABLED=0 GO11MODULE=on go build

FROM alpine:latest
COPY --from=build /src/cmd/ushtrd/ushtrd /ushtrd
EXPOSE 38663
ENTRYPOINT ["/ushtrd"]
