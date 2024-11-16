#!/bin/bash

set -e

LOCKFILE="/tmp/backend_updater.lock"

if [ -f "$LOCKFILE" ]; then
  echo "another instance is running!"
  exit 1
fi

trap "rm -f $LOCKFILE; exit" INT TERM EXIT
touch "$LOCKFILE"

if [ ! -d ".git" ]; then
  echo "${pwd} is not a git-repository dir"
  rm -f "$LOCKFILE"
  exit 1
fi

git fetch origin

LOCAL=$(git rev-parse HEAD)
REMOTE=$(git rev-parse @{u})

if [ "$LOCAL" != "$REMOTE" ]; then
  echo "found update!"

  git stash

  git pull


  if [ -x "./deploy.sh" ]; then
    echo "deploying..."
    ./deploy.sh
  else
    echo "deploy.sh not found/not executable"
  fi
else
fi

rm -f "$LOCKFILE"
trap - INT TERM EXIT
