
# OPA with Python Example

setup

```sh
poetry install
```

execute

```sh
opa run --server authz.rego && python app.py && curl -H "Authorization: YOUR_JWT_TOKEN_HERE"
```
