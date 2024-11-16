#!/bin/bash

set -euo pipefail

DOCKER_COMPOSE_FILE="docker-compose.yml"
BUILD_CONTEXT="."
NEW_TAG=$(date +"%Y%m%d%H%M")

echo "Начало деплоя обновления бэкенда!"

cleanup() {
  if [[ -f "${DOCKER_COMPOSE_FILE}.bak" ]]; then
    mv "${DOCKER_COMPOSE_FILE}.bak" "$DOCKER_COMPOSE_FILE"
  fi
}
trap cleanup EXIT

cp "$DOCKER_COMPOSE_FILE" "${DOCKER_COMPOSE_FILE}.bak"

echo "Пересборка контейнеров..."
docker compose -f $DOCKER_COMPOSE_FILE build

echo "Тегирование образов с новым тегом: $NEW_TAG..."
for IMAGE in $(docker compose -f $DOCKER_COMPOSE_FILE config | grep 'image:' | awk '{print $2}'); do
  NEW_IMAGE="${IMAGE%%:*}:${NEW_TAG}"
  echo "Тегирование $IMAGE -> $NEW_IMAGE"
  docker tag $IMAGE $NEW_IMAGE
  sed -i "s|$IMAGE|$NEW_IMAGE|g" $DOCKER_COMPOSE_FILE
done

echo "Развертывание новых версий..."
docker compose -f $DOCKER_COMPOSE_FILE up -d --remove-orphans

echo "Очистка старых контейнеров..."
docker image prune -f > /dev/null
docker container prune -f > /dev/null

echo "Проверка состояния..."
if docker compose ps | grep -q "Exit"; then
  echo "Ошибка: Один или несколько контейнеров завершились с ошибкой."
  cleanup
  exit 1
fi

rm -f "${DOCKER_COMPOSE_FILE}.bak"

echo "Деплой успешно завершен."
