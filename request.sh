#!/bin/bash -eu

gzip -c ./data/sample.json | curl -X POST --data-binary @- -H "Content-Encoding: gzip" -H "Content-Type: application/json" localhost:18888
