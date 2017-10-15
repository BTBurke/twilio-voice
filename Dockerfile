FROM ubuntu:16.04

RUN apt-get update && apt-get -y upgrade
RUN apt-get install -y apt-transport-https git ca-certificates curl software-properties-common
RUN mkdir -p /voice
RUN mkdir -p /voice/vm

COPY context.mp3 /voice/vm/context.mp3
COPY twilio-voice /voice

EXPOSE 8080
WORKDIR /voice

ENTRYPOINT ["./twilio-voice"]
