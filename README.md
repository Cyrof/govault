# GoVault

A simple, secure CLI-based password vault written in Go. </br>
GoVault stores and retrieves secrets locally using **AES-GCM** encryption and **Argon2id** key derivation

---

## Features

- Secure storage with AES-GCM and Argon2id
- Master password verification
- CLI interface for managing secrets
- Fuzzy search for easier key lookups
- Import/export of encrypted vault data
- Password generator (standalone or integrated into `add`)
- Encrypted local file storage

---

## Prerequisites

- [Go](https://go.dev) 1.20 or higher (if building from source)

---

## Installation

Download the latest binary from the [Releases](https://github.com/Cyrof/govault/releases) page for your OS. </br>
Or build from source:

```bash
go build -o govault ./cmd/govault
```

---

## Usage

First-time run will:

- Prompt you to set a master password
- Generate a salt and hash for verification
- Create an encrypted vault file (`vault.enc`) and metadata file (`meta.json`)

### Command Overview

You can view all available commands and flags by running:

```bash
govault --help
```

To view details for a specific command:

```bash
govault <command> --help
```

Example:

```bash
govault add --help
```

### Example Screenshot

(_Image of govault help here_)

---

## Directory Structure

```java
.
├── assets
├── cmd
│   └── govault
├── internal
│   ├── backup
│   ├── crypto
│   ├── fileIO
│   ├── generator
│   ├── logger
│   ├── model
│   └── vault
└── pkg
    ├── cli
    └── cobraCLI
```

---

## Future Plans

- **Migrate to SQL database** for scalable storage
  - Switch from whole vault encryption to per-row encryption
- **Tagging system** for secrets to enable filtered listing
- Potential future CLI enhancements

---

## License

This project is licensed under the [Apache 2.0](https://github.com/Cyrof/govault/blob/main/LICENSE).
