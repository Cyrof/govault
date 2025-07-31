# Changelog

All notable changes to this project will be documented in this file

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/)

---

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
