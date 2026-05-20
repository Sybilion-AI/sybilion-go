# Changelog

All notable changes to the `go.sybilion.dev/sybilion` Go module are documented here.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Removed

- **`OPERATIONAL_API_BASE_URL`**: no longer read for base URL resolution. Use `SYBILION_API_BASE_URL` or `Options.BaseURL` (breaking for any integration still setting only the old name).

## [0.1.0] - 2026-05-12

### Added

- Initial public release of the **`go.sybilion.dev/sybilion`** Go module.
- `package sybilion` (wrapper) with `New`, `Options`, `Forecasts().Wait`, `ForEachUsagePage`, `ForEachJobsPage`, `AsGenericOpenAPIError`.
- `package sybilionapi` at `go.sybilion.dev/sybilion/api` (OpenAPI 0.1.0 generated client).
- Base URL resolution via `SYBILION_API_BASE_URL` (env), then the compiled default `https://api.sybilion.dev`.

### Changed

- **Renamed** module path from `github.com/Mir-Insight/developers-portal-api-sdk-go` to `go.sybilion.dev/sybilion` (pre-release rename; no prior published version).
- **Renamed** wrapper Go package from `devportalclient` to `sybilion` and the generated package from `devportal` (at `gen/`) to `sybilionapi` (at `api/`).
