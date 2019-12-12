#!/bin/bash

days=()

for dir in ./*/; do
  if [ -d "$dir/Go/" ]; then
    day=$(echo $dir | awk '{match($1,"./(day[0-9]{2})/",a)}END{print a[1]}')
    days+=("$day/Go/$day $day/Go/input.txt")
  fi
done

echo $days

hyperfine "${days[@]}" --warmup 50 --export-csv aoc2019.csv
