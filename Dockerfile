# FROM golang:1.11
FROM rerost/pubsub-proxy:latest

ENV GOPATH /go

ENV APP_ROOT $GOPATH/src/github.com/rerost/pubsub-proxy
WORKDIR /${APP_ROOT}

# Install Dependency
COPY Makefile ${APP_ROOT}
COPY Gopkg.toml ${APP_ROOT}/
COPY Gopkg.lock ${APP_ROOT}

RUN make vendor

# Build Binary
COPY . ${APP_ROOT}/
RUN gex grapi build

RUN ln -sf $APP_ROOT/ /app

CMD ["/app/bin/server"]
