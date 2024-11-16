#!/bin/bash

set -e

LOCKFILE="/tmp/backend_updater.lock"

trap "rm -f $LOCKFILE; exit" INT TERM EXIT
if [ -f "$LOCKFILE" ]; then
  echo "Another instance is running!"
  exit 1
fi
touch "$LOCKFILE"

if ! git rev-parse --show-toplevel > /dev/null 2>&1; then
  echo "$(pwd) is not a git-repository dir"
  rm -f "$LOCKFILE"
  exit 1
fi

if ! git fetch origin; then
  echo "Failed to fetch from origin"
  rm -f "$LOCKFILE"
  exit 1
fi

LOCAL=$(git rev-parse HEAD)
REMOTE=$(git rev-parse @{u})

if [ "$LOCAL" != "$REMOTE" ]; then
  git stash push -q --include-untracked || true
  git pull --rebase

  if [ -x "./deploy.sh" ]; then
    ./deploy.sh
  fi
fi

rm -f "$LOCKFILE"
trap - INT TERM EXIT
