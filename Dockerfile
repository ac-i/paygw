# [1] build
FROM docker.io/library/golang:1.17 AS build
COPY . /src/app
ENV CGO_ENABLED=0
WORKDIR /src/app
RUN go mod tidy && go build -o /app

# [2] run
FROM scratch
COPY --from=build /app /

ENTRYPOINT ["/app"]
LABEL paygw.version="0.0.1"

ENV PAYGW_SERVERS="grpc restgw -rest"
ENV PAYGW_STORAGES="mem -grpc -aql -cql"
ENV GRPC_NETWORK="tcp"
ENV GRPC_ADDRESS=":8090"
ENV RESTGW_ADDRESS=":8080"
# ENV REST_ADDRESS=":8070"
# ENV AQL_HOSTS="localhost"
# ENV AQL_USERNAME="arango"
# ENV AQL_PASSWORD="arango"
# ENV CQL_HOSTS="localhost"
# ENV CQL_USERNAME="scylla"
# ENV CQL_PASSWORD="scylla"

EXPOSE 8080 8090
