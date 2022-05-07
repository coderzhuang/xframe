#!/bin/bash

docker build -t app .
#docker run --rm --name=app --network=work_default -d -p 8080:8080 app
