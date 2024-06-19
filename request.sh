#!/bin/bash -eu

curl -X POST -F "file=@./data.zip" -H "Content-Type: application/zip" localhost:18888
