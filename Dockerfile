FROM golang as builder
WORKDIR /go/src/github.com/LOG-ED/cloud-native-app

# Install and run dep
RUN go get -u github.com/golang/dep/cmd/dep
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only

# Copy the code and compile it
ADD . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /app ./cmd/app/

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app ./
COPY --from=builder /go/src/github.com/LOG-ED/cloud-native-app/tmpl/homepage.html tmpl/homepage.html
CMD ["./app"]