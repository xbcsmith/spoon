#!/bin/bash


mkdir -p temp && cp -v restapi/configure_spoon.go restapi/operations/spoon_api.go temp/
rm -rvf ./restapi ./cmd ./client ./models

echo "Generating Server Code"

swagger generate server -A spoon -P models.Principal -f ./swagger.yml

echo "Generating Client Code"

swagger generate client -A spoon -P models.Principal -f ./swagger.yml


