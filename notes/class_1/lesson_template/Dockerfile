#syntax=docker/dockerfile:1
FROM golang:1.18

ARG SRCDIR="/go/src/github.com/tj10200/golang-class/notes/class_1/lesson_template"
ENV ADDRESS 0.0.0.0
ENV DEBUG false
ENV LESSON_HELLO "HELLO"
ENV LESSON_WORLD "DOCKER!"

RUN mkdir -p $SRCDIR
WORKDIR $SRCDIR
COPY cmd $SRCDIR/cmd
COPY config $SRCDIR/config
COPY pkg $SRCDIR/pkg
COPY go.mod go.sum main.go $SRCDIR

RUN git config --global url."git@github.com:".insteadOf "https://github.com/"
RUN mkdir ~/.ssh
RUN ssh-keyscan -H github.com >> ~/.ssh/known_hosts

RUN --mount=type=ssh cd $SRCDIR && go install

CMD ["lesson_template"]
