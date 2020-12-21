# golang
FROM golang

# build args
ARG PROJECT
ARG TEAM
ARG ENV
ARG GO111MODULE

ENV PROJECT ${PROJECT}
ENV TEAM ${TEAM}
ENV ENV ${ENV}
ENV GO111MODULE ${GO111MODULE}

# check for updates
RUN apt-get update

# install any updates available
RUN apt-get -y upgrade

# create the code location on the container
RUN mkdir -p /go/src/github.com/${TEAM}/${PROJECT}

# copy the local code to the container location
COPY . /go/src/github.com/${TEAM}/${PROJECT}

# move the container's working directory
WORKDIR /go/src/github.com/${TEAM}/${PROJECT}

# make deps
RUN make deps

# build
RUN go build

# run
RUN go install

# clean
RUN make clean