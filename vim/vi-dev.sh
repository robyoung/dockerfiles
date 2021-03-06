#!/usr/bin/env bash

id

set -x

sudo -u docker \
    docker run \
        -ti \
        -u robyoung \
        -e VIM_EXTRA_PLUGINS=1 \
        -e VIMHOME=${HOME}/.vim \
        -e FZF_DEFAULT_COMMAND='rg --files' \
        -w ${PWD} \
        -v ${HOME}/${DEV_DIR}:${HOME}/${DEV_DIR} \
        -v ${HOME}/dev-vim:${HOME}/.vim \
        -v ${HOME}/${DEV_DIR}/personal/dotfiles/vimrc:${HOME}/.vimrc \
        -v ${HOME}/.viminfo:${HOME}/.viminfo \
        -v ${HOME}/${DEV_DIR}/github/junegunn/fzf:${HOME}/.fzf \
        vim \
          vim \
          -i ${HOME}/.vim/viminfo \
          "$@"
