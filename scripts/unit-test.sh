#!/bin/bash

go test -coverprofile=coverage.out ./... | sed ''/FAIL/s//$(printf "\033[31mFAIL\033[0m")/''
go tool cover -func=coverage.out
