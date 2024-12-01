#!/bin/bash
# Runs current day

YEAR=$(TZ=EST date +%Y)
DAY=$(TZ=EST date +%d)
DIR="$YEAR/$DAY"
PREP=bin/prepare
INPUT=input.txt

if [ ! -f "$DIR/$INPUT" ]; then
	if [ ! -x "$PREP" ]; then
		make "$PREP"
	fi

	"$PREP"
fi

cd "$YEAR/$DAY"
exec go run "$DAY.go" "$@"
