# spoon

Silverware things

Yes this is the Spoon to the Silverware!


## Get go-swagger

```
download_url=$(curl -s https://api.github.com/repos/go-swagger/go-swagger/releases/latest | \
  jq -r '.assets[] | select(.name | contains("'"$(uname | tr '[:upper:]' '[:lower:]')"'_amd64")) | .browser_download_url')
curl -o ~/bin/swagger -L'#' "$download_url"
chmod +x ~/bin/swagger
```

## Init

Init is only intended to be run the first time to create a swagger.yml template. No need to run it again as it will overwrite the existing swagger.yml.

```
swagger init spec \
  --title "The spoon application" \
  --description "Spoons are Silverware" \
  --version 0.0.1 \
  --scheme http \
  --consumes application/com.github.xbcsmith.spoon.v1+json \
  --produces application/com.github.xbcsmith.spoon.v1+json
```

## Validate

```
swagger validate swagger.yml
```

2019/05/25 10:44:53

The swagger spec at "swagger.yml" is valid against swagger specification 2.0

## Generate

Not necessary unless you change the swagger.yml
Use the ./scripts/generate.sh as it backs up the config_spoon.go and spoon_api.go file first.

Server

```
swagger generate server -A spoon -P models.Principal -f ./swagger.yml
```

Client

```
swagger generate client -A spoon -P models.Principal -f ./swagger.yml
```

Certs

```
./scripts/gencerts.sh
```

## Test

```
export SPOON_TOKEN="$(echo spoons | sha256sum | awk '{print $1}')
```

```
spoon-server --port 30570 --tls-certificate server.rsa.crt --tls-key server.rsa.key
```

```
curl -k -i -H 'Content-Type: application/com.github.xbcsmith.spoon.v1+json' -H "X-Spoon-Token: $SPOON_TOKEN" http://127.0.0.1:30570/api/v1/spoons
```

```
curl -i -H 'Content-Type: application/com.github.xbcsmith.spoon.v1+json' -H "X-Spoon-Token: $SPOON_TOKEN" http://127.0.0.1:30570/api/v1/spoons -d "{\"description\":\"message $RANDOM\"}"
```

```
curl -i -H 'Content-Type: application/com.github.xbcsmith.spoon.v1+json' -H 'X-Spoon-Token: 12345' http://127.0.0.1:30570/api/v1/spoons/1
```

## Helm

Debug charts

```
export SPOON_TOKEN="$(echo spoons | sha256sum | awk '{print $1}')"
```

```
helm template --set server.x_spoon_token=$SPOON_TOKEN --debug ./helm/spoon/
