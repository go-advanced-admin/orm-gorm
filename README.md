# Go Advanced Admin - GORM Integration

[GORM](https://gorm.io/) integration for the Go Advanced Admin Panel.

[![Go Report Card](https://goreportcard.com/badge/github.com/go-advanced-admin/orm-gorm)](https://goreportcard.com/report/github.com/go-advanced-admin/orm-gorm)
[![Go](https://github.com/go-advanced-admin/orm-gorm/actions/workflows/tests.yml/badge.svg)](https://github.com/go-advanced-admin/orm-gorm/actions/workflows/tests.yml)
[![License: Apache-2.0](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](LICENSE)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/go-advanced-admin/orm-gorm?tab=doc)

This package provides GORM integration for the Go Advanced Admin Panel, allowing you to use GORM as your ORM backend.

## Installation

Add the module to your project by running:

```sh
go get github.com/go-advanced-admin/gorm-integration
```

## Documentation

For detailed documentation on how to use the GORM integration, please visit the [official documentation website](https://goadmin.dev/gprm.html).

## Quick Start

```go
import (
    "github.com/go-advanced-admin/admin"
    "github.com/go-advanced-admin/gorm-integration"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

func main() {
    // Initialize GORM
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // Initialize the ORM integrator
    ormIntegrator := admingorm.NewIntegrator(db)

    // Use ormIntegrator when initializing the admin panel
}
```

For more detailed examples and configuration options, please refer to the [GORM Integration Guide](https://goadmin.dev/gorm.html).

## Contributing

Contributions are always welcome! Please refer to the [Contributing Guidelines](https://github.com/go-advanced-admin/admin/blob/main/CONTRIBUTING.md) in the main repository.

## License

This project is licensed under the Apache 2.0 License - see the [LICENSE](LICENSE) file for details.
