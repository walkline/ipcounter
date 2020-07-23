#!/bin/sh

for (( c=1; c<=1000; c++ ))
do
	ip=`echo "{\"ip\":\"$((RANDOM % 256)).$((RANDOM % 256)).$((RANDOM % 256)).$((RANDOM % 256))\"}"`
	b64=`echo $ip | base64`
  echo "{\"method\":\"POST\",\"url\":\"http://localhost:5000/logs\",\"body\":\"$b64\"}"
done
