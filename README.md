# mysql proxy

### Что это

Это сервис API для Mysql, написанный на Go lang

### Зависимости

Необходимо выполнить:

		 go get github.com/go-sql-driver/mysql
     go get github.com/gorilla/sessions
     go get github.com/gorilla/mux

Зайти в папку *mysqlproxy* и скомпилировать:

		 cd mysqlproxy
		 go build -o ./mysqlproxy *.go


### Как это работает

Скомпилированный файл следует запускть со следующими флагами:

- *mysqlString* - строка подключения к MySQL, например user:pass@/dbname
- *accessToken* - секретный токен, который будет использоваться в запросах к сервису
- *port* - порт. Если не указать, сервис по умолчанию запустится на порту 7000

     mysqlproxy -mysqlString=user:pass@/dbname -accesToken=CHANGEME -port=7000

### Как это использовать

Для авторизации необходимо один раз отправить запрос GET по адресу

     http://localhost:7000/login/*accessToken*

Нужно использовать то значение *accessToken*, которое задано при запуске сервиса

Обращение к MySQL базе делается запросами GET на адрес:

     http://localhost:7000/query/*query_text*

Здесь *query_text* - это запрос к базе, например **select * from table**. Первое слово запроса (select, insert, update или delete должно быть строчными буквами)

На select возвращается результат в JSON, на все остальные (insert, update или delete) - количество измененных или удаленных записей.