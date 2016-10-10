FROM quay.io/eris/build
MAINTAINER Monax Industries <support@monax.io>

#-----------------------------------------------------------------------------
# install eris-cm

ENV REPO $GOPATH/src/github.com/eris-ltd/eris-cm
COPY . $REPO
WORKDIR $REPO/cmd/eris-cm
RUN go build -o $INSTALL_BASE/eris-cm
RUN mkdir /defaults && \
  mv $REPO/account-types /defaults/. && \
  mv $REPO/chain-types /defaults/. && \
  chown --recursive $USER:$USER /defaults

# ----------------------------------------------------------------------------
# mintgen
RUN go get github.com/eris-ltd/mint-client/mintgen
WORKDIR $GOPATH/src/github.com/eris-ltd/mint-client/mintgen
RUN go build -o $INSTALL_BASE/mintgen

#-----------------------------------------------------------------------------
# persist data, set user
RUN rm -rf $GOPATH
RUN chown --recursive $USER:$USER /home/$USER
VOLUME /home/$USER/.eris
WORKDIR /home/$USER/.eris
USER $USER
ENTRYPOINT ["eris-cm"]
