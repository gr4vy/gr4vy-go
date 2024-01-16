#!/bin/bash
rm -rf api
docker run --rm \
  -v ${PWD}:/local openapitools/openapi-generator-cli:v5.1.1 generate \
  -i https://gr4vy.github.io/gr4vy-openapi/openapi.sdks.v1.json \
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

var1='Undefined'
rep1='string'
sed -i '' "s/$var1/$rep1/g" ./api/model_payment_option_context.go

var1='PaymentServiceUpdateFields'
rep1='PaymentServiceRequestFields'
sed -i '' "s/$var1/$rep1/g" ./sdk_payment_services.go

sh replace.sh