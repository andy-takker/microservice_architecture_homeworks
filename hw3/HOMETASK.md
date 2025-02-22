# Основы работы с Kubernetes

## Цель

Научиться создавать минимальный сервис в Kubernetes

## Описание / Пошаговая инструкция выполнения домашнего задания

Шаг 1

Создать минимальный сервис, который

1. Отвечает на порту 8000
2. Имеет http-методы:

   - `GET /health` с ответом `{ "status": "OK" }`
   - `GET /greeting/<name>` с ответом `Hello, <name>!`

Шаг 2

Собрать локально образ приложения в Docker контейнер под архитектуру AMD64 и запушить образ на Docker Hub.

Шаг 3

Написать манифесты для деплоя в k8s для этого сервиса.
Манифесты должны описывать сущности: Deployment, Service, Ingress.
В Deployment могут быть указаны Liveness, Readiness пробы.
Количество реплик должно быть не меньше 2.
Образ контейнера  должен быть указан с Docker Hub.
Хост в Ingress должен быть `arch.homework`.
В итоге после применения манифестов GET запрос на `http://arch.homework/health` должен отдавать `{ "status": "OK" }`.

## Дополнительное задание

В Ingress-е должно быть правило, которое форвардит все запросы с `/otusapp/{student name}*` на сервис с изменением пути, где `{student name}` - это имя студента.

## Рекомендации по форме сдачи

- прописать у себя в `/etc/hosts` хост `arch.homework` с адресом своего minikube (выполните `minikube ip`), чтобы обращение было по имени хоста в запросах, а не IP
- использовать nginx ingress контроллер, установленный через Helm, а не встроенный в minikube:

  ```bash
  kubectl create namespace m
  helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx/
  helm repo update
  helm install nginx ingress-nginx/ingress-nginx --namespace m -f nginx-ingress.yaml
  ```
