FROM golang:1.16-alpine AS build
WORKDIR /go/src/work
ENV CGO_ENABLED=0

# COPY ./cmd/mock-hashicorp-raft-join-server ./cmd/mock-hashicorp-raft-join-server
# COPY ./pkg/log ./pkg/log
# COPY go.* ./

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

COPY . .

# RUN ls .
# RUN cat configs/config.json
# RUN cat secrets/secret.json
# RUN more cmd/mock-hashicorp-raft-join-server/main.go
# RUN pwd

RUN go mod graph | awk '{if ($1 !~ "@") print $2}' | xargs go get

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -gcflags "all=-N -l" -o ./mock-hashicorp-raft-join-server ./cmd/mock-hashicorp-raft-join-server
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /bin/mock-hashicorp-raft-join-server ./cmd/mock-hashicorp-raft-join-server

FROM scratch
COPY --from=build /go/src/work/mock-hashicorp-raft-join-server /mock-hashicorp-raft-join-server
EXPOSE 9000
ENTRYPOINT [ "/mock-hashicorp-raft-join-server" ]
