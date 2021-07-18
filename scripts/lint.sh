#!/bin/bash

golangci-lint run --exclude-use-default=false --out-format tab --enable revive
