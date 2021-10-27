#!/bin/sh

request() {
  curl localhost:"$1"/
}

i=0
while [ $i -ne 500 ]; do
  i=$(($i + 1))
  request "8080" &
  request "8081" &
  request "8082" &
  request "8083" &
  sleep 1
done
