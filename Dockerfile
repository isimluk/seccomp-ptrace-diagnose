FROM ubuntu:latest as builder

RUN : \
    && apt-get update \
    && DEBIAN_FRONTEND=noninteractive apt-get install -y golang ca-certificates git libseccomp-dev

ADD . .

RUN cd diagnose && go build -o diagnose .

FROM ubuntu:latest

COPY --from=builder /diagnose/diagnose /bin/

RUN : \
	&& apt-get update \
  && apt-get install -y strace libcap2-bin \
  && rm -rf /var/lib/apt/lists/*

COPY entrypoint.sh /bin

ENTRYPOINT /bin/entrypoint.sh
