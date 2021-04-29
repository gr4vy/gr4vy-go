#!/bin/bash
rm -rf api
docker run --rm \
  -v ${PWD}:/local openapitools/openapi-generator-cli generate \
  -i https://raw.githubusercontent.com/gr4vy/gr4vy-openapi/sdks/openapi.v1.json \
  -g go \
  --git-user-id gr4vy \
  --git-repo-id gr4vy-go \
  --enable-post-process-file \
  -o /local/api \
  -c /local/.openapi-generator-config.json

rm -rf api/go.mod
rm -rf api/go.sum

sh replace.sh