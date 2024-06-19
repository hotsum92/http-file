#!/bin/bash -eu

curl -X POST -F "file=@./data.zip" -H "Content-Type: multipart/form-data" localhost:18888
