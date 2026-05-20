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
	"os"

	"go.sybilion.dev/sybilion"
)

func main() {
	c := sybilion.New(sybilion.Options{
		Token: os.Getenv("SYBILION_API_TOKEN"),
	})

	me, _, err := c.DefaultAPI().ApiV1MeGet(context.Background()).Execute()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(me.GetUserId(), me.GetAvailableEurCents(), me.GetApiUsageTier())
}
```

The token is an API key (`sk_ops_...`) created in the Developers Portal, or a dashboard session token.

### Base URL

Resolution order:

1. `Options.BaseURL` when non-empty
2. `SYBILION_API_BASE_URL` in the process environment (optional)
3. Compiled default `https://api.sybilion.dev`

The wrapper does not read the API token from the environment; pass `Options.Token` explicitly (often from `SYBILION_API_TOKEN` or your own secret store).

## What's in the box

- `sybilion.Client` — Bearer auth, ergonomic helpers.
- `c.Forecasts().Wait(ctx, jobID, poll)` — polls `GET /api/v1/forecasts/{id}` until settled.
- `sybilion.ForEachUsagePage` / `ForEachJobsPage` — paginated iterators.
- `c.DefaultAPI()` — escape hatch to the OpenAPI-generated client at `go.sybilion.dev/sybilion/api`.
- `sybilion.AsGenericOpenAPIError(err)` — typed error unwrap.

## Documentation

Full guides, feature walkthroughs, API reference, and SDK patterns: **<https://sybilion.dev/docs/>**.

For LLM / codegen contexts the single-file cheat sheet lives at [`LLM_SDK_GUIDE.md`](./LLM_SDK_GUIDE.md).

## Support

- Email — [support@sybilion.com](mailto:support@sybilion.com)
- Slack — [Sybilion Community](https://join.slack.com/t/sybilioncommunity/shared_invite/zt-3y6vx56nk-WJu35eLxkyFQr~Yfko6RjQ)
- Discord — [Sybilion Developers Community](https://discord.gg/CEpEvUB7)

## License

[Apache 2.0](LICENSE).
