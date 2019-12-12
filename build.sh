#!/bin/bash

for dir in ./*/; do
  if [ -d "$dir/Go/" ]; then
    day=$(echo $dir | awk '{match($1,"./(day[0-9]{2})/",a)}END{print a[1]}')
    echo "Building $day ..."
    go build -o "$dir/Go/$day" $dir/Go
  fi
done
