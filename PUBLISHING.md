# Publishing — `go.sybilion.dev/sybilion`

Go modules are not "published" anywhere — they are discovered by the Go module proxy on the first `go get`, which clones the repo at the requested git tag. So a release for this SDK is:

1. The repo is **publicly clonable** at `https://github.com/Sybilion-AI/developers-portal-api-sdk-go`.
2. The vanity URL `go.sybilion.dev/sybilion` resolves via the static page in [`vanity/`](./vanity/) and points the toolchain at that repo.
3. A semver git tag (`vMAJOR.MINOR.PATCH`) on the default branch defines the version.

## One-time setup (do these in order)

### 1. The repo is public

Confirmed already — `gh repo view Sybilion-AI/developers-portal-api-sdk-go --json visibility` returns `"PUBLIC"`. If you ever need to flip back, both vanity URL resolution and `go get` will break.

### 2. (Optional) Rename the repo to `sybilion-go`

Cosmetic but matches the package name. **Settings -> Rename repository -> `sybilion-go`**. GitHub keeps the old URL as a permanent redirect, so the vanity meta tag continues to work even before you re-deploy. After the rename, update [`vanity/index.html`](./vanity/index.html) and re-deploy the vanity page so its `<meta name="go-import">` and `<meta name="go-source">` tags point at `https://github.com/Sybilion-AI/sybilion-go`.

### 3. Add DNS for `go.sybilion.dev`

In Cloudflare DNS (or wherever `sybilion.dev` is managed), add a `CNAME` record for `go` pointing at the host you'll deploy the vanity page to. If you use Cloudflare Pages (recommended below), the dashboard adds this record automatically once you set the custom domain.

### 4. Deploy the vanity page

See the deployment guide in [`vanity/README.md`](./vanity/README.md). Quick version with Cloudflare Pages:

```bash
npx wrangler pages deploy vanity --project-name=sybilion-vanity
```

Then in the Cloudflare dashboard set `go.sybilion.dev` as the project's custom domain.

### 5. Verify resolution

```bash
curl 'https://go.sybilion.dev/sybilion?go-get=1' | grep go-import
# Expected:
# <meta name="go-import" content="go.sybilion.dev/sybilion git https://github.com/Sybilion-AI/developers-portal-api-sdk-go" />

GOPROXY=direct go install go.sybilion.dev/sybilion@latest 2>&1 | head
# Should attempt a fetch (will report a real error only if no tag exists yet).
```

## Per-release flow

```bash
# 1. Bump and document the version
$EDITOR CHANGELOG.md           # move [Unreleased] -> [0.1.0] - YYYY-MM-DD

# 2. Make sure local validation is green
make verify                    # regenerate + drift check + tests

# 3. Commit, tag, push
git add CHANGELOG.md
git commit -m "release: v0.1.0"
git tag v0.1.0
git push origin master v0.1.0
```

The tag push triggers `release.yml` which validates the tree and creates a GitHub Release with notes generated from commits. **`pkg.go.dev` discovers the new version automatically** the first time anyone runs `go get go.sybilion.dev/sybilion@v0.1.0`. Force-warm the cache by running it yourself once.

## Sanity checks after the first release

```bash
# In a clean Go module, fetch the published version.
mkdir /tmp/sybilion-go-test && cd /tmp/sybilion-go-test
go mod init sandbox
go get go.sybilion.dev/sybilion@v0.1.0
cat <<'EOF' > main.go
package main

import (
	"fmt"

	"go.sybilion.dev/sybilion"
)

func main() {
	fmt.Println(sybilion.DefaultPublicAPIBaseURL)
}
EOF
go run .
cd .. && rm -rf sybilion-go-test
```

Within a few minutes the version page should appear at <https://pkg.go.dev/go.sybilion.dev/sybilion>. If it does not, visit that URL to trigger discovery.

## Retracting a bad release

If you ship a broken version, add a `retract` directive to `go.mod` in a follow-up release rather than deleting the tag (deleted tags are a Go module proxy nightmare).

```go
module go.sybilion.dev/sybilion

go 1.25.10

retract v0.1.0  // critical bug in Forecasts().Wait, fixed in v0.1.1
```

Tag and push `v0.1.1`. Existing pinned consumers keep working; new `go get` resolutions skip the retracted version.
