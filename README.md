# volt-demo-clis

The **monorepo + mixed** demo for [volt](https://github.com/khanakia/voltkit): two CLIs with independent version streams, a nested importable library, and shared internal code — one repo, one pipeline.

| Directory | Kind | Tag stream |
|---|---|---|
| `cmd/notes` | CLI | `notes/vX.Y.Z` — tag is the **binary name** |
| `cmd/snap` | CLI | `snap/vX.Y.Z` — independent of notes |
| `pkg/textutil` | library (own `go.mod`) | `pkg/textutil/vX.Y.Z` — tag is the **directory path**, proxy-resolvable |
| `internal/shared` | never released | — |

Released with `volt release ./cmd/notes v0.1.0` etc. Doubles as a regression fixture for volt.
