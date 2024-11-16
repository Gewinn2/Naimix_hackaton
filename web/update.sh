#!/bin/bash

set -e

LOCKFILE="/tmp/auto_update.lock"

if [ -f "$LOCKFILE" ]; then
  echo "Скрипт уже выполняется. Пропуск запуска."
  exit 1
fi

trap "rm -f $LOCKFILE; exit" INT TERM EXIT
touch "$LOCKFILE"

if [ ! -d ".git" ]; then
  echo "Ошибка: Текущая директория не является Git-репозиторием."
  rm -f "$LOCKFILE"
  exit 1
fi

echo "Проверка обновлений в репозитории..."
git fetch origin


LOCAL=$(git rev-parse HEAD)
REMOTE=$(git rev-parse @{u})

if [ "$LOCAL" != "$REMOTE" ]; then
  echo "Найдены обновления. Обновление репозитория..."

  git stash

  git pull


  if [ -x "./deploy.sh" ]; then
    echo "Запуск deploy.sh..."
    ./deploy.sh
  else
    echo "Скрипт deploy.sh не найден или не является исполняемым."
  fi
else
  echo "Обновлений не найдено."
fi

rm -f "$LOCKFILE"
trap - INT TERM EXIT

echo "Скрипт завершён."
