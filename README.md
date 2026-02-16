# go-analyzers-golangcilint

[golangci-lint v2 module plugin](https://golangci-lint.run/docs/plugins/module-plugins/) for [go-analyzers](https://github.com/albertocavalcante/go-analyzers).

Registers the following analyzers:

- **`makecopy`** — detects `make` + `copy` that can be replaced with `slices.Clone`
- **`searchmigrate`** — detects `sort.Search` that can be replaced with `slices.BinarySearch`
- **`clampcheck`** — detects if/else-if clamping that can be replaced with `min`/`max`

## Usage

1. Create `.custom-gcl.yml` in your project:

```yaml
version: v2.9.0
plugins:
  - module: 'github.com/albertocavalcante/go-analyzers-golangcilint'
    version: v0.1.0
```

2. Build custom binary:

```bash
golangci-lint custom
```

3. Add to your golangci-lint config:

```toml
[linters.settings.custom.go-analyzers]
type = "module"
description = "Custom analyzers for modern Go idioms"
```

4. Run:

```bash
./custom-gcl run ./...
```

## License

[MIT](LICENSE)
