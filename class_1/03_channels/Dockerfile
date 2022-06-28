#syntax=docker/dockerfile:1
FROM golang:1.18

ARG SRCDIR="/go/src/github.com/tj10200/golang-class/class_1/03_channels"

RUN mkdir -p $SRCDIR
WORKDIR $SRCDIR

# Copy the source into the destination
COPY cmd $SRCDIR/cmd
COPY config $SRCDIR/config
COPY pkg $SRCDIR/pkg
COPY go.mod go.sum main.go $SRCDIR

# Allow github access for private repos during install
#  The actual ssh key is passed in at build time
RUN git config --global url."git@github.com:".insteadOf "https://github.com/"
RUN mkdir ~/.ssh
RUN ssh-keyscan -H github.com >> ~/.ssh/known_hosts

# Pulls any required 3rd parties and installs them as dependencies.
RUN --mount=type=ssh cd $SRCDIR && go install

CMD ["go", "run", "main.go", "lesson"]

