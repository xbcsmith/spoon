#!/bin/bash

TEMPDIR=$(mktemp -d -p ./temp/)

cp -v restapi/configure_spoon.go restapi/operations/spoon_api.go $TEMPDIR/

rm -rvf ./restapi ./cmd ./client ./models

echo "Generating Server Code"

swagger generate server -A spoon -P models.Principal -f ./swagger.yml

echo "Generating Client Code"

swagger generate client -A spoon -P models.Principal -f ./swagger.yml

diff -rupN $TEMPDIR/spoon_api.go restapi/operations/spoon_api.go > $TEMPDIR/spoon_api.diff

diff -rupN $TEMPDIR/configure_spoon.api restapi/configure_spoon.go > $TEMPDIR/configure_spoon.diff


