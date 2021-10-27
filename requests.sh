#!/bin/sh

request() {
  curl localhost:8080/"$1"
}

i=0
while [ $i -ne 500 ]; do
  i=$(($i + 1))
  request "zap" &
  request "logrus" &
  request "zerolog" &
  request "stdlib" &
  sleep 1
done
