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
spoon-server --port 30570 --tls-certificate server.rsa.crt --tls-key server.rsa.key
```

```
curl -k -i -H 'Content-Type: application/com.github.xbcsmith.spoon.v1+json' -H 'X-Spoon-Token: 12345' http://127.0.0.1:30570/api/v1/spoons
```
