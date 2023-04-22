
# OPA with Golang Concert Example

setup

```sh
go mod tidy && go mod vendor
```

execute

```sh
go run main.go <bracelet_color>
# colors are orange, blue, purple
```

execute tests

```sh
opa test -v policy.rego policy_test.rego
```
