#!/bin/bash

set -e
exec > /root/.hackaton/backend/backend_updater.log 2>&1
set -x

LOCKFILE="/tmp/backend_updater.lock"
trap "rm -f $LOCKFILE; exit" INT TERM EXIT

if [ -f "$LOCKFILE" ]; then
  echo "Another instance is running!"
  exit 1
fi

touch "$LOCKFILE"

# Переход в директорию проекта
cd /root/.hackaton/backend || exit 1

# Проверка на git репозиторий
if ! /usr/bin/git rev-parse --show-toplevel > /dev/null 2>&1; then
  echo "$(pwd) is not inside a git repository"
  exit 1
fi

# Обновление кода
if ! /usr/bin/git fetch origin; then
  echo "Error: Failed to fetch from origin"
  exit 1
fi

LOCAL=$(/usr/bin/git rev-parse HEAD)
REMOTE=$(/usr/bin/git rev-parse @{u})

if [ "$LOCAL" != "$REMOTE" ]; then
  /usr/bin/git stash push -q --include-untracked || true
  /usr/bin/git pull --rebase

  if [ -f "./deploy.sh" ] && [ -x "./deploy.sh" ]; then
    ./deploy.sh
  else
    echo "Warning: deploy.sh is missing or not executable" >&2
  fi
fi

rm -f "$LOCKFILE"
trap - INT TERM EXIT
