# appMagic

![GitHub last commit](https://img.shields.io/github/last-commit/burmatovdd/appMagic?style=flat-square)
![GitHub top language](https://img.shields.io/github/languages/top/burmatovdd/appMagic?style=flat-square)

# :memo: Задача
> Реализовать сервис обработки истории цены gas в сети ethereum.
> 
# :crayon: Подзадачи 
- Подсчитать сколько было потрачено gas помесячно
- Рассчитать среднюю цену gas за день
- Вычислить частотное распределение цены по часам(за весь период)
- Выясниить сколько заплатили за весь период

# :pushpin: Проблемы с которыми столкнулся и как их решил
- Большой объем данных трудно обрабатывать. 
- Решение: Разбил данные по временам, стало несколько массивов с которыми удобно работать
- Частотное распределение для каждого часа работало некорректно из-за того, что все цены отличались хотябы одним знаком. 
- Решение: Округлил цены до целого, чтобы наглядно было видно, что сервис работает корректно. 

# :rocket: Запуск проекта
> Запуск сервиса (запускать из корневой папки)
> 
> `go run cmd/server/main.go`

# :question: Как пользоваться сервисом
> При запуске, проект стартует на 8080 порту
>
> Необходимо послать на сервер запрос 
>
> `curl localhost:8080/api/getData`

# :card_index_dividers: Стек технологий
 - gin-gonic
 - viper config


