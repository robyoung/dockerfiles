FROM debian:buster-slim

ARG VIM_USERNAME=default
ARG VIM_USERID=1001
RUN groupadd -g ${VIM_USERID} ${VIM_USERNAME} && \
  useradd -d /home/${VIM_USERNAME} -m -u ${VIM_USERID} -g ${VIM_USERID} ${VIM_USERNAME}

RUN apt-get update && apt-get -y install vim git python3 python3-pip editorconfig
COPY fzf /usr/bin/
COPY rg /usr/bin/

USER ${VIM_USERNAME}
