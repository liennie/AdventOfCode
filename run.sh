#!/bin/bash
# Runs current day

set -e

YEAR=$(TZ=EST date +%Y)
MONTH=$(TZ=EST date +%m)
DAY=$(TZ=EST date +%d)

if [[ $MONTH != 12 || $DAY > 25 ]]; then
	echo "🟡 Today is not an advent day"
	exit 1
fi

DIR="$YEAR/$DAY"
PREP=bin/prepare
INPUT=input.txt

if [ ! -f "$DIR/$INPUT" ]; then
	if [ ! -x "$PREP" ]; then
		make "$PREP"
	fi

	"$PREP"

	if $(which code >/dev/null); then
		code "$DIR/$DAY.go" "$DIR/test.txt"
	fi
fi

cd "$DIR"
exec go run "$DAY.go" "$@"
