#!/bin/bash

for ((i = 0; i < 10000; i++)); do
    curl "http://127.0.0.1:8001/"
    curl "http://127.0.0.1:8001/hello"
    curl "http://127.0.0.1:8001/business"
done
