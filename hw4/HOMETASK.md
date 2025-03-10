# Инфраструктурные паттерны

## Цель

Создать простейший RESTful CRUD.

## Описание / Пошаговая инструкция выполнения домашнего задания

- Сделать простейший RESTful CRUD по созданию, удалению, просмотру и обновлению пользователей. [Пример API](https://app.swaggerhub.com/apis/otus55/users/1.0.0).
- Добавить базу данных для приложения.
- Конфигурация приложения должна храниться в Configmaps
- Доступы к БД должны храниться в Secrets.
- Первоначальные миграции должны быть оформлены в качестве Job-ы, если это требуется.
- Ingress-ы должн ытакже вести на URL `arch.homework` (как и в прошлом задании)

На выходе должны быть предоставлены:

1. Ссылка на директорию в GitHub, где находится директория с манифестами Kubernetes
2. Инструкция по запуску приложения
   - команда установки БД из helm, вместе с файлом `values.yaml`
   - команда применения первоначальных миграций
   - команда `kubectl apply -f`, которая запускает в правильном порядке манифесты Kubernetes
3. Postman коллекция, в которой будут представлены примеры запросов к сервису на создание, получение, изменение и удаление пользователя. Важно: в Postman коллекции использовать базовый URL - `arch.homework`
4. Проверить корректность работы приложения, используя созданную коллекцию через `newman run ./postman_collection.json` и приложить скриншот/вывод исполнения корректной работы.

## Задание со звездочкой

Добавить шаблонизацию приложения в Helm чартах.
