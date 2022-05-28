FROM golang:1.16-alpine AS build
WORKDIR /go/src/work
ENV CGO_ENABLED=0

COPY ./cmd/hermes ./cmd/hermes
COPY ./pkg ./pkg
COPY go.* ./

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN go mod graph | awk '{if ($1 !~ "@") print $2}' | xargs go get

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -gcflags "all=-N -l" -o ./hermes ./cmd/hermes

FROM scratch
COPY --from=build /go/src/work/hermes /hermes
EXPOSE 8000
EXPOSE 9000
EXPOSE 10000
ENTRYPOINT [ "/hermes" ]
