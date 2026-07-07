# Sybilion Go SDK

Official Go client for the [Sybilion API](https://sybilion.dev/docs/).

```bash
go get go.sybilion.dev/sybilion@latest
```

Requires Go 1.25+.

## Quick use

```go
package main

import (
	"context"
	"fmt"
	"log"

	"go.sybilion.dev/sybilion"
)

func main() {
	// Token read from SYBILION_API_TOKEN env var automatically
	c := sybilion.New(sybilion.Options{})

	me, err := c.Me(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(me.GetUserId(), me.GetAvailableEurCents(), me.GetApiUsageTier())
}
```

Or pass the token explicitly:

```go
c := sybilion.New(sybilion.Options{Token: "sk_ops_..."})
```

The token is an API key (`sk_ops_...`) created in the Developers Portal, or a dashboard session token.

### Base URL

Resolution order:

1. `Options.BaseURL` when non-empty
2. `SYBILION_API_BASE_URL` in the process environment (optional)
3. Compiled default `https://api.sybilion.dev`

## What's in the box

**Account**
- `c.Me(ctx)` — authenticated account info (balance, tier).

**Catalog**
- `c.ListCategories(ctx)` / `c.ListRegions(ctx)` — available thematic categories and geographic regions.

**Forecasts**
- `c.SubmitForecast(ctx, req)` — submit an async forecast job. Set `req.AuxTimeseries` (1–10 series, each keyed on the same dates as `Timeseries`) to supply your own forecast drivers.
- `c.GetForecast(ctx, id)` — poll status of a forecast job.
- `c.GetForecastArtifact(ctx, id, name)` — download a forecast artifact by name.
- `c.Forecasts().Wait(ctx, jobID, poll)` — polls until the job is settled or context is done.

**Drivers**
- `c.GetDrivers(ctx, req)` — drivers ranked by explanatory power (synchronous, billed).

**Alerts**
- `c.GetAlerts(ctx, req)` — anomaly alerts for a timeseries (synchronous, billed).

**Jobs & Usage**
- `c.ForEachJobsPage(ctx, ...)` / `c.ForEachUsagePage(ctx, ...)` — paginated iterators.

**Escape hatch**
- `c.DefaultAPI()` — the OpenAPI-generated client at `go.sybilion.dev/sybilion/api`.

## Documentation

Full guides, feature walkthroughs, API reference, and SDK patterns: **<https://sybilion.dev/docs/>**.

For LLM / codegen contexts the single-file cheat sheet lives at [`LLM_SDK_GUIDE.md`](./LLM_SDK_GUIDE.md).

## Support

- Email — [support@sybilion.com](mailto:support@sybilion.com)
- Slack — [Sybilion Community](https://join.slack.com/t/sybilioncommunity/shared_invite/zt-3y6vx56nk-WJu35eLxkyFQr~Yfko6RjQ)
- Discord — [Sybilion Developers Community](https://discord.gg/KMDyXBdQ8c)

## License

[Apache 2.0](LICENSE).
