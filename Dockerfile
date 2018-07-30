### STAGE 1: Build ###
FROM golang:1.10 AS builder

# Set the working directory to the app directory
WORKDIR /go/src/dewey

# ADD all source code to WORKDIR
ADD . /go/src/dewey

# get dependant Go packages
RUN go get github.com/olivere/elastic
RUN go get github.com/confluentinc/confluent-kafka-go/kafka
RUN go get github.com/sirupsen/logrus

# compile as Linux executable
RUN cd /go/src/dewey && CGO_ENABLED=0 GOOS=linux go build -o eventTracker

### STAGE 2: Setup ###
FROM scratch 

# set working directory
WORKDIR /root/

# copy the binary from builder and place in WORKDIR
COPY --from=builder /go/src/dewey .

# run the binary
CMD ["./eventTracker"]