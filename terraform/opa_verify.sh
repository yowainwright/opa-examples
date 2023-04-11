#!/bin/bash

# Generate Terraform plan JSON
terraform init
terraform plan -out=tfplan.binary
terraform show -json tfplan.binary > tfplan.json

# Check OPA policy
opa eval --format pretty --data policy.rego --input tfplan.json 'data.main.deny'
