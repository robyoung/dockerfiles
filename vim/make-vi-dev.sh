#!/bin/bash

go build -o vi-dev vi-dev.go \
  && sudo setcap cap_setgid=ep vi-dev
