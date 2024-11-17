FROM ubuntu:latest

WORKDIR /usr/bin
RUN apt update
RUN apt install -y wget
RUN wget https://github.com/shaka-project/shaka-packager/releases/download/v3.0.4/packager-linux-arm64
RUN mv packager-linux-arm64 packager
RUN chmod +x ./packager


