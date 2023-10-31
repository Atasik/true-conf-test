Здравствуйте, %USERNAME%!

Вам предстоит выполнить рефакторинг небольшого приложения на Go (200 строк).

Приложение представляет собой API по работе с сущностью User, где хранилищем выступает файл json.

Ограничения:
- Хранилищем должен оставаться файл в json формате.
- Структура пользователя не должна быть уменьшена.
- Приложение не должно потерять существующую функциональность. 

Мы понимаем, что пределу совершенства нет и ожидаем, что объем рефакторинга вы определяете на свое усмотрение.  

После того как вы выполните задание, вы так же можете написать, как бы улучшили проект в перспективе текстом.

Что следует знать:
- В будущем это приложение ожидает увеличение количества функций и сущностей. 
- Вопрос авторизации умышленно опущен, о нем не стоит беспокоиться.
- API еще не выпущено, вы в праве скорректировать интерфейс / форматы ответов.

Работа должна быть оформлена на Github, а все изменения выполнены в отдельном(ых) коммитах.

Удачи!

# Что можно улучшить
- Добавить DTO структуры на каждый уровень приложения, чтобы компоненты приложения были более независимыми
- В перспективе добавить контейнеризацию для разворачивания приложения и оркестратор этими контейрнерами (docker+kuber)
- Добавить тесты, пайплайн, логгирование, валидацю полей структур
- Прописать отдельные тайм-ауты для специфических запросов, я посчитал это избыточным, поэтому просто поставил общие timeout-ы на сервере