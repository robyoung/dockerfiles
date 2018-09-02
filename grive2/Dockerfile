FROM alpine:latest as builder
LABEL maintainer "Rob Young <rob@robyoung.digital>"

RUN apk update
RUN apk add git
RUN git clone https://github.com/vitalif/grive2.git /usr/src/grive2

RUN apk add alpine-sdk cmake \
      libgcrypt-dev curl-dev expat-dev yajl-dev boost-dev

RUN cd /usr/src/grive2 && \
      mkdir build && \
      cd build && \
      cmake .. && \
      make -j4 

FROM alpine:latest

RUN apk update && apk add libgcrypt libcurl expat yajl boost boost-program_options shadow

COPY --from=builder /usr/src/grive2/build/grive/grive /usr/local/bin/grive

ARG GRIVE_USERNAME=default
ARG GRIVE_USERID=1001
RUN groupadd -g ${VIM_USERID} ${VIM_USERNAME} && \
  useradd -d /home/${VIM_USERNAME} -m -u ${VIM_USERID} -g ${VIM_USERID} ${VIM_USERNAME}

USER ${VIM_USERNAME}

WORKDIR /mnt
