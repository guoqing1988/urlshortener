#!/usr/bin/env bash

output=`curl -s -v -sX POST 'localhost:8080/shorten' -d http://www.google.com`

echo ""
echo "### Got short URL: $output"
echo "### Requesting shortened URL:"
echo ""

curl -s -v $output
