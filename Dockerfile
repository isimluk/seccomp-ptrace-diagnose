FROM ubuntu:latest

RUN : \
	&& apt-get update \
  && apt-get install -y strace \
  && rm -rf /var/lib/apt/lists/*

COPY entrypoint.sh /bin

ENTRYPOINT /bin/entrypoint.sh
