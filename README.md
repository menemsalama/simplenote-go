# simplenote-go
dummy clone of simplenote, and POC for working with postgresql in golang

## Running Locally

Makefile `run` command will generate the proper docs module to serve it at `/swagger/index.html` (or `/swagger/*`) and run go run {main.go}

To generate swagger docs and start serving

```sh
$ make run
```

Required .env vars

1. SECRET=somesecret
2. DB_URL=postgresqlurl
