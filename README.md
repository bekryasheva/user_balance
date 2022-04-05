# User-balance 
Тестовое задание User-balance – это микросервис для работы с балансом пользователей.

### Для запуска приложения:

В docker-compose.yaml указать значение переменной окружения CURRENCY_API_ACCESS_KEY (ключ API для использования <a href="https://currencyapi.com/">currencyapi</a>).

```
docker-compose up
```
### Методы API
`/deposit`

Метод начисления средств на баланс. Принимает id пользователя и сколько средств зачислить.

Пример:
```
curl -i -X POST "http://localhost:8080/deposit" \
    -H "Content-Type: application/json" \
    -d '{"id":1,"amount":1000}' 
```   
Возможные ответы:
``` 
HTTP/1.1 200 OK
Date: Mon, 04 Apr 2022 21:02:30 GMT
Content-Length: 0
```
``` 
HTTP/1.1 400 Bad Request
Content-Type: application/json; charset=UTF-8
Date: Tue, 05 Apr 2022 10:19:13 GMT
Content-Length: 38

{"message":"input validation failed"}
``` 

`/withdrawal`

Метод списания средств с баланса. Принимает id пользователя и сколько средств списать.

Пример:
```
curl -i -X POST "http://localhost:8080/withdrawal" \
    -H "Content-Type: application/json" \
    -d '{"id":1,"amount":10}'
```    
Возможные ответы:
```
HTTP/1.1 200 OK
Date: Tue, 05 Apr 2022 10:06:23 GMT
Content-Length: 0 
```
```
HTTP/1.1 412 Precondition Failed
Content-Type: application/json; charset=UTF-8
Date: Tue, 05 Apr 2022 10:10:24 GMT
Content-Length: 33

{"message":"insufficient funds"}
```
```
HTTP/1.1 412 Precondition Failed
Content-Type: application/json; charset=UTF-8
Date: Tue, 05 Apr 2022 10:11:19 GMT
Content-Length: 32

{"message":"user is not found"}
```
```
HTTP/1.1 400 Bad Request
Content-Type: application/json; charset=UTF-8
Date: Tue, 05 Apr 2022 10:15:05 GMT
Content-Length: 38

{"message":"input validation failed"}
```
`/transfer`

Метод перевода средств от пользователя к пользователю. Принимает id пользователя с которого нужно списать средства, id пользователя которому должны зачислить средства, а также сумму.

Пример:
```
curl -i -X POST "http://localhost:8080/transfer" \
    -H "Content-Type: application/json" \
    -d '{"fid":1, "tid":2, "amount":10}' 
```
Возможные ответы:
```
HTTP/1.1 200 OK
Date: Tue, 05 Apr 2022 10:06:45 GMT
Content-Length: 0
```
```
HTTP/1.1 412 Precondition Failed
Content-Type: application/json; charset=UTF-8
Date: Tue, 05 Apr 2022 10:20:55 GMT
Content-Length: 33

{"message":"insufficient funds"}
```
```
HTTP/1.1 412 Precondition Failed
Content-Type: application/json; charset=UTF-8
Date: Tue, 05 Apr 2022 10:21:52 GMT
Content-Length: 32

{"message":"user is not found"}
```
```
HTTP/1.1 400 Bad Request
Content-Type: application/json; charset=UTF-8
Date: Tue, 05 Apr 2022 10:22:24 GMT
Content-Length: 38

{"message":"input validation failed"}
```

`/balance/:id`

Метод получения текущего баланса пользователя. Принимает id пользователя. Если присутствует доп.параметр currency, то баланс пользователя конвертируется с рубля на указанную валюту. Данные по текущему курсу валют берутся из <a href="https://currencyapi.com/">currencyapi</a>.

Пример:
```
curl -i -X GET "http://localhost:8080/balance/1?currency=USD"
```
Возможные ответы:
```
HTTP/1.1 200 OK
Content-Type: application/json; charset=UTF-8
Date: Tue, 05 Apr 2022 10:07:15 GMT
Content-Length: 44

{"id":1,"balance":11.7012,"currency":"USD"}
```
```
HTTP/1.1 412 Precondition Failed
Content-Type: application/json; charset=UTF-8
Date: Tue, 05 Apr 2022 10:25:39 GMT
Content-Length: 32

{"message":"user is not found"}
```
```
HTTP/1.1 400 Bad Request
Content-Type: application/json; charset=UTF-8
Date: Tue, 05 Apr 2022 10:26:06 GMT
Content-Length: 38

{"message":"input validation failed"}
```

`/history/:id`

Метод получения списка транзакций. Сортировка по сумме и дате с помощью order_by и sort, пагинация с помощью limit, offset. Для получения списка транзакций также будет достаточно указать лишь order_by.

Пример:
```
curl -i -X GET "http://localhost:8080/history/1?order_by=date&sort=asc&limit=3&offset=1"
```
Возможные ответы:
```
HTTP/1.1 200 OK
Content-Type: application/json; charset=UTF-8
Date: Tue, 05 Apr 2022 10:07:40 GMT
Content-Length: 346

[{"id":2,"user_id":1,"type_of_transaction":"withdrawal","date":"2022-04-05T10:06:23.17396Z","amount":10,"from_id":{"Int64":0,"Valid":false},"to_id":{"Int64":0,"Valid":false}},{"id":4,"user_id":1,"type_of_transaction":"transfer","date":"2022-04-05T10:06:45.66704Z","amount":10,"from_id":{"Int64":1,"Valid":true},"to_id":{"Int64":2,"Valid":true}}]
```
```
HTTP/1.1 412 Precondition Failed
Content-Type: application/json; charset=UTF-8
Date: Tue, 05 Apr 2022 10:59:38 GMT
Content-Length: 32

{"message":"user is not found"}
```
```
HTTP/1.1 400 Bad Request
Content-Type: application/json; charset=UTF-8
Date: Tue, 05 Apr 2022 10:27:31 GMT
Content-Length: 38

{"message":"input validation failed"}
```


