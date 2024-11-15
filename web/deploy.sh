#!/bin/bash

set -e

APP_NAME="naimix-frontend"
DOCKERFILE="Dockerfile"
BUILD_CONTEXT="."
NEW_TAG=$(date +"%Y%m%d%H%M")

echo "Начинаю деплой фронтенда!"

echo "Сборка Docker-образа..."
docker build -f $DOCKERFILE -t $APP_NAME:latest $BUILD_CONTEXT

echo "Тегирование нового образа: $NEW_TAG..."
docker tag $APP_NAME:latest $APP_NAME:$NEW_TAG

echo "Остановка и удаление старого контейнера..."
if docker ps | grep -q $APP_NAME; then
  docker stop $APP_NAME || true
  docker rm $APP_NAME || true
fi

echo "Запуск нового контейнера..."
docker run -d --name $APP_NAME $APP_NAME:$NEW_TAG

echo "Деплой успешно завершен!"
