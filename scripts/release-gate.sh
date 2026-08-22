#!/bin/sh
# release-gate.sh — pre_release hook demo: an arbitrary project-owned gate.
# volt runs this AFTER tests, BEFORE the tag; non-zero exit aborts the
# release with nothing created. RELEASE_BLOCKED=1 simulates a failing gate.
if [ "$RELEASE_BLOCKED" = "1" ]; then
  echo "gate: releases are blocked (RELEASE_BLOCKED=1) — refusing $VOLT_TAG" >&2
  exit 1
fi
echo "gate: ok — $VOLT_TAG ($VOLT_KIND at $VOLT_DIR)"
