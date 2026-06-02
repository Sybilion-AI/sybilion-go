# Changelog

All notable changes to the `go.sybilion.dev/sybilion` Go module are documented here.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).


## [Unreleased]

## [0.1.3] - 2026-05-28

### Added

- **`Client.GetAlerts()`**: new `POST /api/v1/alerts` endpoint. New exported types `AlertsRequest`, `AlertItem`, and `NewsItem`.
- **`Client.Me()`**: ergonomic wrapper for `GET /api/v1/me`.
- **`Client.ListCategories()` / `Client.ListRegions()`**: catalog endpoint wrappers.
- **`Client.SubmitForecast()` / `Client.GetForecast()` / `Client.GetForecastArtifact()`**: individual forecast wrappers (previously only `Forecasts().Wait` existed).
- **`Client.GetDrivers()`**: wrapper for `POST /api/v1/drivers`.
- **`EnvSybilionAPIToken`** constant (`"SYBILION_API_TOKEN"`).
- **Token from env**: `Options.Token` is now optional; when empty, `SYBILION_API_TOKEN` env var is read automatically.
- **API error messages**: non-2xx responses now parse the `error` field from the JSON response body and surface it as a plain `error`.

### Changed

- **`WaitForecast` is now a method**: `func WaitForecast(ctx, api, jobID, poll)` → `(*Client).WaitForecast(ctx, jobID, poll)`. Breaking for callers using the package-level function.
- **`ForEachUsagePage` / `ForEachJobsPage` are now methods**: `(*Client).ForEachUsagePage` / `(*Client).ForEachJobsPage`. Breaking for callers using the package-level functions.
- **`Forecasts().Wait` delegates to `Client.WaitForecast`**.

### Removed

- **`Client.Raw()`**: use `Client.DefaultAPI()` to access the underlying generated client. Breaking for callers using `Raw()`.
- **`APIError` type and `AsGenericOpenAPIError` function**: replaced by transparent `parseAPIError` error wrapping. Breaking for callers inspecting error types directly.

## [0.1.2] - 2026-05-28

### Changed

- Updated Discord community invite link.

## [0.1.1] - 2026-05-20

### Changed

- **Forecast horizon**: split into **`soft_horizon`** and **`hard_horizon`** (Go fields **`SoftHorizon`** and **`HardHorizon`**).

### Removed

- **`OPERATIONAL_API_BASE_URL`**: no longer read for base URL resolution. Use `SYBILION_API_BASE_URL` or `Options.BaseURL` (breaking for any integration still setting only the old name).

## [0.1.0] - 2026-05-12

### Added

- Initial public release of the **`go.sybilion.dev/sybilion`** Go module.
- `package sybilion` (wrapper) with `New`, `Options`, `Forecasts().Wait`, `ForEachUsagePage`, `ForEachJobsPage`, `AsGenericOpenAPIError`.
- `package sybilionapi` at `go.sybilion.dev/sybilion/api` (OpenAPI 0.1.0 generated client).
- Base URL resolution via `SYBILION_API_BASE_URL` (env), then the compiled default `https://api.sybilion.dev`.

### Changed

- **Renamed** module path from `github.com/Sybilion-AI/developers-portal-api-sdk-go` to `go.sybilion.dev/sybilion` (pre-release rename; no prior published version).
- **Renamed** wrapper Go package from `devportalclient` to `sybilion` and the generated package from `devportal` (at `gen/`) to `sybilionapi` (at `api/`).
