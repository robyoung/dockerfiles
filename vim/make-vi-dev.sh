#!/bin/bash

go build -o vi-dev vi-dev.go \
  && sudo chown root:root vi-dev \
  && sudo chmod u+s vi-dev
