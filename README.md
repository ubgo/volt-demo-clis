# volt-demo-clis

The **monorepo + mixed** demo for [volt](https://github.com/khanakia/voltkit): two CLIs with independent version streams, a nested importable library, and shared internal code — one repo, one pipeline.

| Directory | Kind | Tag stream |
|---|---|---|
| `.` (root module) | library, AWS-shaped (go.mod at root, packages in subdirs) | bare `vX.Y.Z` |
| `cmd/notes` | CLI | `notes/vX.Y.Z` — tag is the **binary name** |
| `cmd/snap` | CLI | `snap/vX.Y.Z` — independent of notes |
| `pkg/textutil` | library (own `go.mod`) | `pkg/textutil/vX.Y.Z` — tag is the **directory path**, proxy-resolvable |
| `pkg/mathx` | library (own `go.mod`) | `pkg/mathx/vX.Y.Z` — same shape, second published stream |
| `pkg/testutil` | **`internal: true`** | none, ever — tested by `volt ci`, hidden from `volt status`, refused by `volt release` |
| `internal/shared` | Go-`internal` package of the root module | — (releases with the root) |

`volt status` here shows every releasable stream and only those:

```
.                library  v0.1.0    up to date
cmd/notes        cli      v0.2.0    up to date
cmd/snap         cli      v0.1.0    up to date
pkg/mathx        library  v0.1.0    up to date
pkg/textutil     library  v0.1.0    up to date
```

The internal marker is one committed file:

```yaml
# pkg/testutil/.volt.yml
internal: true    # applies to this directory and everything under it
```

Released with `volt release ./cmd/notes v0.1.0` etc. Doubles as a regression fixture for volt.
