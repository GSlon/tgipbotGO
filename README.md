# telegram bot для получения информации об ip

## запуск (docker)
docker-compose run

## запуск (без контейнера)
* make migrate -> отдельно миграция;
* make run -> запукс (с миграцией)

## docker_env:
* DB_PASSWORD -> password for connect to db
* BOT_TOKEN -> telegram bot token

* DB_PASSWORD, POSTGRES_PASSWORD -> same for user and db

Меню юзера
- Проверить айпи через бота.
- Вывести историю запросов.

Админ меню
- Массовая рассылка всем пользователям когда либо писавшим боту, сообщения которое отправит админ.
- Добавление и удаление новых админов.
- Вывод всех айпи что проверял конкретный пользователь.

