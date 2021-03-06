FROM golang:1.16-alpine AS build
WORKDIR /go/src/work
ENV CGO_ENABLED=0

# COPY ./cmd/http-log-client ./cmd/http-log-client
# COPY ./pkg/log ./pkg/log
# COPY go.* ./

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

COPY . .

# RUN ls .
# RUN cat configs/config.json
# RUN cat secrets/secret.json
# RUN more cmd/http-log-client/main.go
# RUN pwd

RUN go mod graph | awk '{if ($1 !~ "@") print $2}' | xargs go get

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go get -ldflags "-s -w -extldflags '-static'" github.com/go-delve/delve/cmd/dlv@master
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -gcflags "all=-N -l" -o ./http-log-client ./cmd/http-log-client
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /bin/http-log-client ./cmd/http-log-client


FROM scratch
COPY --from=build /go/bin/dlv /dlv
COPY --from=build /go/src/work/http-log-client /http-log-client
EXPOSE 2345
ENTRYPOINT [ "/dlv" ]
