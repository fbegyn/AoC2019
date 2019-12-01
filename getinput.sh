#!/bin/sh

DAY=$1
COOKIE=$AoCCookie

curl \
  -fSL -o ./input.txt \
  -H "cookie:$COOKIE" \
  https://adventofcode.com/2019/day/$DAY/input
