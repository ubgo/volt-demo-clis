// Package shared is code both CLIs use and nothing ever releases — it exists
// so the demo exercises volt ci's fallback: a change here maps to no released
// directory, so the safe answer is "gate everything".
package shared

// Banner prefixes every demo CLI's output so shared-code changes are visible
// in both binaries.
func Banner(app string) string { return "[demo] " + app }
