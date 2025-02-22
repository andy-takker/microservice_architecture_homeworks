# Решение

Я реализовал два http-метода:

- `GET /health` с ответом статуса
- `GET /otusapp/:name/*` с переадресацией на `/greeting/:name` и возвратом ответа `Hello, :name!` 

Чтобы применить все манифесты выполните:

```bash
kubectl apply -f ./manifests
```

Если возникнут проблемы с применением конфигурации ingress, то необходимо выполнить:

```bash
kubectl edit cm -n ingress-nginx ingress-nginx-controller
```

и добавить:

```yaml
data:
  allow-snippet-annotations: "true"
```

## Тестирование

Для проверки работоспособности сервиса выполните:

```bash
curl arch.homework/health
curl arch.homework/health/

curl arch.homework/otusapp/petr/
curl arch.homework/otusapp/petr/something
```
