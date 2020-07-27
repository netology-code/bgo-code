## "Справочные" команды

```shell
docker help

Usage:  docker [OPTIONS] COMMAND

A self-sufficient runtime for containers

Options:
  ...
  -v, --version            Print version information and quit

Management Commands:
  config      Manage Docker configs
  container   Manage containers
  image       Manage images
  system      Manage Docker
  ...

Run 'docker COMMAND --help' for more information on a command.
```

```shell
docker system info

Client:
 Debug Mode: false

Server:
 Containers: 0
  Running: 0
  Paused: 0
  Stopped: 0
 Images: 0
 Server Version: 19.03.8
 ...
```

## Запуск первого контейнера

Скачиваем образ:
```shell
docker image pull hello-world

Using default tag: latest
latest: Pulling from library/hello-world
0e03bdcc26d7: Pull complete
Digest: sha256:d58e752213a51785838f9eed2b7a498ffa1cb3aa7f946dda11af39286c3db9a9
Status: Downloaded newer image for hello-world:latest
docker.io/library/hello-world:latest
```

Создаём контейнер:
```shell
docker container create --name first hello-world
8792116c417895c3d58f9d32e105718dd8aea4faf647b23238763626c9b35207
```

Запускаем контейнер:
```shell
docker container start first
first
```

## Опции запуска

Запускаем контейнер PostgreSQL с переменными окружения (но без binding'а портов):
```shell
docker container run -d \
    --name bankdb \
    -e POSTGRES_PASSWORD=pass \
    -e POSTGRES_USER=app \
    -e POSTGRES_DB=db \
    postgres
Unable to find image 'postgres:latest' locally
latest: Pulling from library/postgres
...
Digest: sha256:9ba6355d27ba9cd0acda1e28afaae4a5b7b2301bbbdc91794dcfca95ab08d2ef
Status: Downloaded newer image for postgres:latest
7e43c41c189c4f783c66191e86b976873a8c2edbffcba973c14d32145de9a690
```

Останавливаем контейнер:
```shell
docker container stop bankdb
```

Удаляем контейнер:
```shell
docker container rm bankdb
```

Запускаем контейнер с binding'ом портов (при этом создаётся новый контейнер из образа, старый мы удалили):
```shell
docker container run -d \
    --name bankdb \
    -e POSTGRES_PASSWORD=pass \
    -e POSTGRES_USER=app \
    -e POSTGRES_DB=db \
    -p 5432:5432 \
    postgres
```

Выполняем команду внутри контейнера (psql):
```shell
docker container exec -it postgres psql -h db -U postgres
```
