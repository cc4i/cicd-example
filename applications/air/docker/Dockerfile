FROM golang:alpine as builder

ENV GO111MODULE=on
ARG ver=0.0.0


# Install git
RUN apk update && apk add git

COPY . $GOPATH/src/mypackage/myapp/
RUN sed -i -e "s/-1.0.0/$ver/g" $GOPATH/src/mypackage/myapp/airdex.go
RUN rm $GOPATH/src/mypackage/myapp/airdex_api.go
WORKDIR $GOPATH/src/mypackage/myapp/


#get dependancies
#you can also use dep
RUN go get -d -v

#build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/airdex

# STEP 2 build a small image
# start from scratch
FROM scratch

# Copy our static executable
COPY --from=builder /go/bin/airdex /go/bin/airdex
COPY ./ChinaCityList.json /ChinaCityList.json
#EXPOSE 8080
CMD ["/go/bin/airdex"]