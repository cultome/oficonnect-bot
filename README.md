# OfiConnect Bot

Bot to check and automatically register into OfiConnect events

# Usage

To check and register events automatically use the `daemon`

```bash
$ go build -o bin/daemon cmd/daemon/main.go
# Get this token after sign-in in https://v2.oficonnect.omdai.org/oficial/inicio
$ set -x OFICONNECT_TOKEN <oficonnect_token>
$ ./bin/daemon <oficonnect_id>
```

To add an event to the exclusion list (not registering automatically), use the `manager`

```bash
$ go build -o bin/manager cmd/manager/main.go
$ ./bin/manager --exclude <eventID>
```
