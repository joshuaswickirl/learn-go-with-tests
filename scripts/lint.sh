#!/bin/bash

golangci-lint run \
    --exclude-use-default=true \
    --out-format tab \
    --config=.golangci.yml
