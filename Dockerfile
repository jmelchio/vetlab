FROM golang:1.25.10-alpine3.23 AS test-build-stage
LABEL authors="joris.melchior@gmail.com"

COPY ./go.mod /workspace/go.mod
COPY ./go.sum /workspace/go.sum
COPY ./api /workspace/api
COPY ./cmd /workspace/cmd
COPY ./model /workspace/model
COPY ./repository /workspace/repository
COPY ./service /workspace/service
COPY ./tools /workspace/tools
COPY ./vetlabcmd /workspace/vetlabcmd

RUN cd /workspace; set -e; \
    go mod download; \
    go get github.com/onsi/ginkgo/v2/...; \
    go get github.com/onsi/gomega/...; \
    go install github.com/onsi/ginkgo/v2/ginkgo; \
    go mod tidy -compat=1.17; \
    ginkgo api/... model/... service/...;\
    go build -o /opt/vetlab/vetlab ./cmd/vetlab

WORKDIR /workspace
ENTRYPOINT ["ginkgo", "repository/..."]

FROM alpine:3.23.0 AS image-stage
LABEL authors="joris.melchior@gmail.com"

COPY --from=test-build-stage /opt/vetlab/vetlab /opt/vetlab/vetlab

EXPOSE 8080

ENTRYPOINT ["/opt/vetlab/vetlab"]