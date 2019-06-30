FROM golang as build-stage
WORKDIR /go
# RUN go get ...
RUN go get github.com/valyala/fasthttp
RUN go get github.com/pkg/errors
RUN go get github.com/rs/xid

# Copy the server code into the container
COPY . /go


RUN go build

# Production
FROM golang as production-stage
WORKDIR /go
COPY --from=build-stage /go/go /go/go
COPY --from=build-stage /go/files /go/files
# COPY --from=build-stage /go/TestCerts /go
EXPOSE 443
EXPOSE 80
ENTRYPOINt ["./go"]