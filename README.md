# go-analyzers-gcl

[![CI](https://github.com/albertocavalcante/go-analyzers-gcl/actions/workflows/ci.yml/badge.svg)](https://github.com/albertocavalcante/go-analyzers-gcl/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/albertocavalcante/go-analyzers-gcl.svg)](https://pkg.go.dev/github.com/albertocavalcante/go-analyzers-gcl)
[![Go Report Card](https://goreportcard.com/badge/github.com/albertocavalcante/go-analyzers-gcl)](https://goreportcard.com/report/github.com/albertocavalcante/go-analyzers-gcl)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

[golangci-lint v2](https://golangci-lint.run/) [module plugin](https://golangci-lint.run/docs/plugins/module-plugins/) for [go-analyzers](https://github.com/albertocavalcante/go-analyzers) -- custom Go static analyzers for modern Go 1.25+ idioms.

## Analyzers

| Analyzer | Detects | Suggests |
|---|---|---|
| `makecopy` | `make([]T, len(s))` + `copy(dst, s)` | `slices.Clone(s)` |
| `searchmigrate` | `sort.Search(n, func(i int) bool { ... })` | `slices.BinarySearch(s, v)` |
| `clampcheck` | `if x < lo { x = lo } else if x > hi { x = hi }` | `min(max(x, lo), hi)` |

These fill gaps left by the official [modernize](https://pkg.go.dev/golang.org/x/tools/go/analysis/passes/modernize) suite. See the [go-analyzers README](https://github.com/albertocavalcante/go-analyzers#why-these-analyzers) for details.

## Usage

### 1. Create `.custom-gcl.yml`

```yaml
version: v2.9.0
plugins:
  - module: 'github.com/albertocavalcante/go-analyzers-gcl'
    version: v0.1.0
```

### 2. Build the custom binary

```bash
golangci-lint custom
```

### 3. Configure the linters

Add to `.golangci.yml`:

```yaml
version: "2"
linters:
  default: none
  enable:
    - go-analyzers
  settings:
    custom:
      go-analyzers:
        type: module
        description: "Custom analyzers for modern Go idioms"
```

### 4. Run

```bash
./custom-gcl run ./...
```

## License

[MIT](LICENSE)
