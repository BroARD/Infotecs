# Go RESTful API для стажировки в Infotecs. 

Раннее я писал для себя проект по этому же заданию, но с дополнениями, в виде кэша(Redis) и немного другой структурой проекта.
Тот проект слегка перегружен и более распаложен к масштабированию. 
Вот ссылка: https://github.com/BroARD/TransactionAPI

Данное приложение представляет из себя REST API с 3 EndPoint`ами:

* Send на создание транзакции и перевода суммы с одного кошелька на другой
* GetLast для получения последних N транзакций
* GetBalance для получения баланса кошелька по UUID
 
## Для запуска приложение содержит Docker и Docker-compose файлы.
(Так как база данных запускается из контейнера)

## Этапы установи приложения:

```shell
# Клоновироть репозиторий:

git clonehttps://github.com/BroARD/Infotecs.git

# Перейти в директорию:

cd Infotecs

# Запустить сборку Docker-compose:

docker-compose -f docker-compose.yml up -d
```

Данный Rest API будет запущен по адрессу: `http://127.0.0.1:8080`. Он содержит такие Эндпоинты как:

* `POST /api/send`: Принимает JSON файл формата и проводит транзакцию:
```
{
    "from" : "",
    "to" : "",
    "amount" : 10
}
```
* `GET /api/transactions?count=N`: Выводит N последних транзакций
* `GET /api/wallet/:wallet_id/balance`: Возвращает данные кошелька по Wallet_ID


To use the starter kit as a starting point of a real project whose package name is `github.com/abc/xyz`, do a global 
replacement of the string `github.com/qiangxue/go-rest-api` in all of project files with the string `github.com/abc/xyz`.


## Project Layout

The starter kit uses the following project layout:
 
```
├───cmd                  
│   └───api              main файл для запуска
├───config               конфиг приложения
├───internal             ядро приложения
│   ├───entity           все сущности для БД
│   ├───middleware       проверки входящих данных
│   ├───server           запуск сервера
│   ├───transaction      всё связанное с транзакциями
│   │   └───dto          образы для входящих данных
│   └───wallet           всё связанное с кошельками
├───logs                 логги
└───pkg                  папка с БД и логгером
    ├───db               всё для БД
    └───logging          всё для Логгера
```  
