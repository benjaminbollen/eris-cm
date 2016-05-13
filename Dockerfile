FROM quay.io/eris/base
MAINTAINER Eris Industries <support@erisindustries.com>

#-----------------------------------------------------------------------------
# dependencies

RUN apt-get update && \
  apt-get install -y --no-install-recommends \
    libgmp3-dev jq && \
  rm -rf /var/lib/apt/lists/*

#-----------------------------------------------------------------------------
# install eris-cm

ENV REPO $GOPATH/src/github.com/eris-ltd/eris-cm
COPY . $REPO
WORKDIR $REPO/cmd/eris-cm
ENV GO15VENDOREXPERIMENT 1
RUN go build -o /usr/local/bin/eris-cm
RUN chown --recursive $USER:$USER $REPO

# ----------------------------------------------------------------------------
# mintgen

RUN go get github.com/eris-ltd/mint-client/mintgen

#-----------------------------------------------------------------------------
# persist data, set user
RUN chown --recursive $USER:$USER /home/$USER
VOLUME /home/$USER/.eris
WORKDIR /home/$USER/.eris
USER $USER
ENTRYPOINT ["eris-cm" ]
