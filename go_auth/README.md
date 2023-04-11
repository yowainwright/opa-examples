
# OPA with Golang Auth Example

setup

```sh
go mod tidy && go mod vendor
```

execute

```sh
opa run --server authz.rego && go run main.go && curl -H "Authorization: YOUR_JWT_TOKEN_HERE"
# opa run --server authz.rego && go run main.go && curl -H "Authorization: eyJhbGciOiJIUzI1NiJ9.eyJSb2xlIjoiQWRtaW4iLCJJc3N1ZXIiOiJJc3N1ZXIiLCJVc2VybmFtZSI6IkphdmFJblVzZSIsImV4cCI6MTY4MTE3Njg0MywiaWF0IjoxNjgxMTc2ODQzfQ.msOjNrAFNOr6nFKb84oV75gVMSl5Bg3CAEHXGfYYCBM"
```
