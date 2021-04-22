#!/bin/bash

curl https://raw.githubusercontent.com/gr4vy/gr4vy-openapi/sdks/openapi.v1.json > openapi.v1.json
oapi-codegen openapi.v1.json > api/gr4vy.gen.go
