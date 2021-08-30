#
# daring-cyclops-worker
#
# docker build . -t daring-cyclops-worker:1
#
# docker run --rm -it daring-cyclops-worker:1 /bin/sh
# docker run --rm daring-cyclops-worker:1
#
FROM golang:1.16-alpine
LABEL build_date="2021-09-04"
LABEL description="daring-cyclops-worker"
LABEL maintainer="guycole@gmail.com"
#
RUN apk add git
#
WORKDIR /app
#
COPY go.mod .
COPY go.sum .
RUN go mod download
#
COPY *.go ./
#
#RUN go mod init 
RUN go build -o /worker

# This is for documentation purposes only.
# To actually open the port, runtime parameters
# must be supplied to the docker command.
EXPOSE 8080

# (Optional) environment variable that our dockerised
# application can make use of. The value of environment
# variables can also be set via parameters supplied
# to the docker command on the command line.
#ENV HTTP_PORT=8081

# Run
CMD [ "/worker" ]
