#!/bin/bash
printf "\nRegenerating swagger doc\n\n"
go install github.com/swaggo/swag/cmd/swag@latest
time swag init -d cmd/myapp,internal/access/http,pkg/common
printf "\nDone.\n\n"