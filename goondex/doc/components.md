@startuml Goondex Containers
  !includeurl https://raw.githubusercontent.com/RicardoNiepel/C4-PlantUML/master/C4_Component.puml

  Container(script, "Start script/cron", "cron, shell", "Запуск по расписанию или при старте приложения")
  Container(webclient, "Web Client", "HTML, javascript")

  Boundary(goondex, "Goondex", "bin/exe") {
    Component(api, "API", "net/http", "Основное API для поиска")
    Component(engine, "Engine", "pkg/engine", "Поисковый движок")
    Component(crawler, "Crawler", "pkg/crawler", "Сканирование сайтов")
    Component(index, "Index", "pkg/index", "Индексированные страницы")
    Component(storage, "Webpages Storage", "pkg/webpages/storage", "Хранение страниц")

  }

  ContainerDb(rdb, "Relation Database", "PostgreSQL")
  ContainerDb(keydb, "Key-Value Database", "Redis")

  Rel(script, crawler, "Запуск сканирования сайтов", "crawler.Scan")
  Rel(webclient, api, "Отправляет запрос на поиск", "JSON/HTTP")
  Rel(api, engine, "Поиск страниц по токену", "engine.Search")
  Rel(engine, index, "Поиск индентификаторов страниц", "index.Search")
  Rel(engine, storage, "Поиск страниц по индентификатору", "storage.Find")
  Rel(crawler, index, "Добавление страниц в индекс", "index.Add")
  Rel(crawler, storage, "Добавление страниц", "storage.Insert")
  Rel(index, keydb, "Поиск индентификаторов по токену")
  Rel(index, rdb, "Хранение индекса")
  Rel(storage, rdb, "Хранение страниц")
@enduml
