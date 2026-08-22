#!/bin/sh
# announce.sh — post_release hook demo: runs after publish+verify.
# Writes evidence OUTSIDE the repo (a hook writing into the working tree
# would dirty it and break --from-tag republishes).
LOG="${ANNOUNCE_LOG:-/tmp/volt-announce.log}"
echo "released $VOLT_TAG version=$VOLT_VERSION binary=$VOLT_BINARY dist=$VOLT_DIST" >> "$LOG"
echo "announce: logged $VOLT_TAG to $LOG"
