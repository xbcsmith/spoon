#!/bin/bash

set -e

[[ -n $SPOON_PORT ]] || export SPOON_PORT="30570"
[[ -n $SPOON_HOST ]] || export SPOON_HOST="0.0.0.0"
[[ -n $SPOON_TLS_CERT ]] || export SPOON_TLS_CERT="/usr/local/bin/server.rsa.crt"
[[ -n $SPOON_TLS_KEY ]] || export SPOON_TLS_KEY="/usr/local/bin/server.rsa.key"

# if given a command, run that
if [[ -n "$1" ]]; then
    exec "$@"
fi

# Readiness
rm -rf /tmp/ready && echo "{\"alive\" : true}" > /tmp/ready

/usr/local/bin/spoon-server --host "${SPOON_HOST}" --port "${SPOON_PORT}" --tls-certificate "${SPOON_TLS_CERT}" --tls-key "${SPOON_TLS_KEY}"
