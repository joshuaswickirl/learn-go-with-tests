#!/bin/bash

go test -bench=. -coverprofile=coverage.out ./... | \
    sed ''/PASS/s//$(printf "\033[32mPASS\033[0m")/'' | \
    sed ''/FAIL/s//$(printf "\033[31mFAIL\033[0m")/''
echo
go tool cover -func=coverage.out
