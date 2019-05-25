#!/bin/bash


SUBJECT='/C=US/ST=North Carolina/L=Cary/O=Spoon Company/CN=spoons.io'
# ECDSA recommendation key ≥ secp384r1
# List ECDSA the supported curves (openssl ecparam -list_curves)
openssl req -x509 -nodes -newkey ec:secp384r1 -keyout server.ecdsa.key -out server.ecdsa.crt -days 3650 -subj "${SUBJECT}"

# RSA recommendation key ≥ 2048-bit
openssl req -x509 -nodes -newkey rsa:2048 -keyout server.rsa.key -out server.rsa.crt -days 3650 -subj "${SUBJECT}"

# ECDSA recommendation key ≥ secp384r1
# List ECDSA the supported curves (openssl ecparam -list_curves)
openssl req -x509 -nodes -newkey ec:secp384r1 -keyout client.ecdsa.key -out client.ecdsa.crt -days 3650 -subj "${SUBJECT}"

# RSA recommendation key ≥ 2048-bit
openssl req -x509 -nodes -newkey rsa:2048 -keyout client.rsa.key -out client.rsa.crt -days 3650 -subj "${SUBJECT}"
