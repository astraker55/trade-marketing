# trade-marketing

Микросервис представляет из себя API для работы со статистикой. 
После выполнения команды docker-compose up запускается на 9200 порту.
Доступные эндпоинты:

1) / - Возвращает приветствие 

2) _/savestat_ - сохраняет полученный json-объект с информацией о событии и вовзращает сообщение с этим объектом.
  Пример тела запроса:
  **{"date": "2003-10-01",
     "views": 0, 
     "clicks": 0, 
     "cost": 1238}**
  Ответ на запрос: Event was added: {Date: 2003-10-01 00: 00: 00 +0000 UTC Views: 0 Clicks: 0 Cost: 1238 Cpc: 0 Cpm: 0}
  __Важно__: Добавить при обращении можно только одно событие, поля CPC и CPM будут пустые.

3) _/getstat?from=<some_date>&to=<some_date>&sort=<some_field>_ - получение информации по событиям, поля **from,  to** обязательные
для сортировки можно выбрать **values, clicks, cost, date**. По умолчанию используется поле **date**. Пример обращения:
_http://192.168.99.103:9200/getstat?from=1999-10-01&to=2031-11-01&sort=clicks__

При получении статистики рассчитываются значения cpc и cpm.

4) _/dropstat_ - выполняется сброс информации

Запуск проводился на Windows 10
