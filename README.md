# OfiConnect Bot

Bot to check and automatically register into OfiConnect events

# Usage

```bash
$ go build -o bin/daemon cmd/daemon/main.go
# Get this token after sign-in in https://v2.oficonnect.omdai.org/oficial/inicio
$ set -x OFICONNECT_TOKEN <oficonnect_token>
$ ./bin/daemon <oficonnect_id>
```
