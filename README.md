# 🔨 Knox

[![Test & Build](https://github.com/zeiss/knox/actions/workflows/main.yml/badge.svg)](https://github.com/zeiss/knox/actions/workflows/main.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/zeiss/knox.svg)](https://pkg.go.dev/github.com/zeiss/knox)
[![Go Report Card](https://goreportcard.com/badge/github.com/zeiss/knox)](https://goreportcard.com/report/github.com/zeiss/knox)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Taylor Swift](https://img.shields.io/badge/secured%20by-taylor%20swift-brightgreen.svg)](https://twitter.com/SwiftOnSecurity)

Knox is a simple and secure Terraform backend. 

## Features

- **Simple**: Knox is a simple and secure Terraform backend.
- **Secure**: Knox uses a secure and encrypted storage backend.
- **Fast**: Knox is fast and lightweight.
- **Versioned**: Knox supports versioning of the Terraform state.

Knox has a team-based management of the Terraform state. It is designed to be used in a multi-team environment where each team has its own workspace.

## Terraform

To use Knox as a Terraform backend, you need to configure the backend in your Terraform configuration file.

The url contains the `team/project/environment` name. The `team` is the team name, the `project` is the project name, and the `environment` is the environment name.

```hcl
terraform {
  backend "http" {
    username       = "super"
    password       = "secret"
    address        = "http://localhost:8084/client/zeiss/demo/dev/state"
    lock_address   = "http://localhost:8084/client/zeiss/demo/dev/lock"
    unlock_address = "http://localhost:8084/client/zeiss/demo/dev/unlock"
    lock_method    = "POST"
    unlock_method  = "POST"
  }
}
```

## Helm Chart

There is a Helm chart available for Knox. You can find it in the [helm/charts](/helm/charts) directory.

:warning: **Please note that the Helm chart is still in development and should not be used in production.**

Knox requires a PostgreSQL database to store the state. [CockroachDB](https://www.cockroachlabs.com/) is recommended for production use.

```bash
helm repo add knox https://zeiss.github.io/knox/helm/charts
helm repo update
helm search repo knox
```

## License

[Apache 2.0](/LICENSE)