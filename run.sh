#!/bin/bash
# Runs current day

YEAR=$(TZ=EST date +%Y)
DAY=$(TZ=EST date +%d)

cd "$YEAR/$DAY"
exec go run $DAY.go "$@"
