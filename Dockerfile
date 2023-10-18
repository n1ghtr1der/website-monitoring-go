FROM ubuntu:22.04

ARG DIR
ARG STATS
ARG REPO


WORKDIR /app
COPY . .
RUN apt-get update
RUN apt install -y python2 && apt install -y gnuplot && apt install -y git

RUN mkdir ${STATS}
RUN git clone ${REPO}

RUN cd /app/gitstats && ./gitstats /${DIR} /${DIR}/${STATS}