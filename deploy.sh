#!/bin/bash

set -e

DOCKER_COMPOSE_FILE="docker-compose.yml"
BUILD_CONTEXT="."
NEW_TAG=$(date +"%Y%m%d%H%M")

echo "Начало деплоя обновления бэкенда!"

echo "Пересборка контейнеров..."
docker-compose -f $DOCKER_COMPOSE_FILE build

echo "Тегирование: $NEW_TAG..."
for IMAGE in $(docker-compose -f $DOCKER_COMPOSE_FILE config | grep 'image:' | awk '{print $2}'); do
  NEW_IMAGE="${IMAGE%%:*}:${NEW_TAG}"
  docker tag $IMAGE $NEW_IMAGE
  sed -i "s|$IMAGE|$NEW_IMAGE|g" $DOCKER_COMPOSE_FILE
done

echo "Развертывание..."
docker-compose -f $DOCKER_COMPOSE_FILE up -d --build --remove-orphans

echo "Очистка старых образов и контейнеров..."
docker image prune -f
docker container prune -f

echo "Проверка..."
if docker-compose ps | grep -q "Exit"; then
  echo "Ошибка: Один или несколько контейнеров не работают."
  exit 1
fi

echo "Деплой успешно завершен."
