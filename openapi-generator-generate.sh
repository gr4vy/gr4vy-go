#!/bin/bash
rm -rf api
docker run --rm \
  -v ${PWD}:/local openapitools/openapi-generator-cli:v5.1.1 generate \
  -i https://raw.githubusercontent.com/gr4vy/gr4vy-openapi/sdks/openapi.v1.json \
  -g go \
  --git-user-id gr4vy \
  --git-repo-id gr4vy-go \
  --enable-post-process-file \
  -o /local/api \
  -c /local/.openapi-generator-config.json

rm -rf api/go.mod
rm -rf api/go.sum

var1='var environment ENVIRONMENT = "production"'
rep1='var environment string = "production"'
sed -i '' "s/$var1/$rep1/g" ./api/*.go

var1='Undefined'
rep1='ThreeDSecureDataV2'
sed -i '' "s/$var1/$rep1/g" ./api/model_transaction_request.go