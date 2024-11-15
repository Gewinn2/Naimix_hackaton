#!/bin/bash

set -e

BUILD_CONTEXT="."
DOCKERFILE="./Dockerfile"
IMAGE_NAME="naimix-frontend"
NEW_TAG=$(date +"%m%d%H%M")

echo "Начало деплоя обновления фронтенда!"

echo "Пересборка контейнера..."
docker build -f $DOCKERFILE -t $IMAGE_NAME:latest $BUILD_CONTEXT

echo "Тегирование: $NEW_TAG..."
docker tag $IMAGE_NAME:latest $IMAGE_NAME:$NEW_TAG

echo "Остановка текущего контейнера..."
if docker ps | grep -q $IMAGE_NAME; then
  docker stop $(docker ps -q --filter ancestor=$IMAGE_NAME:latest)
fi

echo "Удаление старого контейнера..."
if docker ps -a | grep -q $IMAGE_NAME; then
  docker rm $(docker ps -a -q --filter ancestor=$IMAGE_NAME:latest)
fi

echo "Развертывание нового контейнера..."
docker run -d --name "${IMAGE_NAME}_${NEW_TAG}" -p 3000:3000 $IMAGE_NAME:$NEW_TAG

echo "Очистка старых образов..."
docker image prune -f

echo "Проверка..."
if ! docker ps | grep -q "${IMAGE_NAME}_${NEW_TAG}"; then
  echo "Ошибка: Контейнер не запущен."
  exit 1
fi

echo "Деплой успешно завершен."
