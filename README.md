# GoVault 
A simple CLI-based password vault written Go. Securely stores and retrieves secrets locally using AES encryption and Argon2-based key derivation.

---
## Features
- Secure storage using AES-GCM and Argon2id
- Master password verification
- CLI interface for adding and retrieving secrets
- Encrypted local file storage

---

## Prerequisites
- [Go](https://go.dev) 1.20 or higher

---

## Usage
> :warning: This project is currently in development. A CI pipeline will be added later to automatically generate binaries via GitHub Actions.

### Build
```bash
go build -o govault ./cmd/govault
```

### Run
```bash
go run ./cmd/govault/main.go
```

### Commands
```bash
./govault add -key <key> -value <value>
./govault get -key <key
```
If no command is passed: 
```bash
./govault
```
A usage guide will be displayed.

---

## Directory Structure
```java
.
├── cmd
│   └── govault
├── go.mod
├── go.sum
├── internal
│   ├── crypto
│   ├── fileIO
│   ├── model
│   └── vault
├── LICENSE
├── pkg
│   └── cli
└── README.md
```
---

## Future Plans
- Implement base CLI with flags
- Convert to [Cobra](https://github.com/spf13/cobra) for richer CLI UX
- Add `view` or `list` command to show stored keys
- Add `purge` command to completely reset vault and metadata (e.g., remove `meta.json` and `vault.enc` files)
- Move secret storage to SQLite (encrypted)
- Optional: support password hints or recovery tokens
- GitHub Actions CI pipeline to build & release executables

---

## License
This project is licensed under the [Apache 2.0](https://github.com/Cyrof/govault/blob/main/LICENSE).
