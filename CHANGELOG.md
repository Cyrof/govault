# Changelog

All notable changes to this project will be documented in this file

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/)

---

## [v0.1.2] - 2025-08-08

### Added

- **Fuzzy Search**: Implemented fuzzy search in the CLI to help users find keys even if the exact name is uncertain. ([#53](https://github.com/Cyrof/govault/pull/53))
- **Import / Export**: Added encrypted import and export functionality for vault data, allowing secure transfer between devices. ([#52](https://github.com/Cyrof/govault/pull/52))
- **Password Generator**: Added a CLI `generate` flag to create random passwords (default: 16 characters, mixed case and numbers, optional symbols). Integrated into the `add` command via a `--gen` subflag. ([#51](https://github.com/Cyrof/govault/pull/51))
- **Edit Feature**: Added the ability to update the password of an existing secret without recreating it. ([#49](https://github.com/Cyrof/govault/pull/49))
- **Delete Feature**: Added a `delete` command to remove stored keys. ([#48](https://github.com/Cyrof/govault/pull/48))

### Changed

- Refactored password generation logic for integration with both the standalone `generate` command and the `add` command. ([#51](https://github.com/Cyrof/govault/pull/51))
- Improved CLI prompts and user interaction flow for new features.
- Minor refactoring to support new import/export and search logic. ([#52](https://github.com/Cyrof/govault/pull/52))

### Fixed

- Corrected flawed vault file creation logic on first run - vault is now created even if no secrets are added initially. ([#48](https://github.com/Cyrof/govault/pull/48))
- Fixed various spelling mistakes in CLI messages.

### Removed

- _No removals in this release._

## [v0.1.1] - 2025-07-31

### Added

- Integrated structured loggin using `zap` with `lumberjack` for log file rotation.
- Configured environment-based logging for `dev` and `prod` modes.
- Enhanced CLI output with colored formatting using the `fatih/color` package for better UX.

### Changed

- Replaced the original CLI implementation using Go's standard `flag` package with `cobra` for more extensible and user-friendly command handling.

### Removed

- Nothing removed in this release

## [v0.1.0] - 2025-07-28

### Added

- `Purge` command allows users to completely reset govault by deleting `meta.json` and `vault.enc` files
- `List` command allows users to list out all the keys stored in the vault without the values
- Added new GitHub CI to automate version bump, changelogs templating, auto build, and release for golang binaries

### Changed

- Updated the `README` to include the guides for the new features

### Removed

- Nothing removed in this release
