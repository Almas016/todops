
# Как запустить
1. `git clone github.com/Almas016/todops`
2. `cd todops`
3. Создать `.env`, после чего заполнить следующие поля
	```
	DB_HOST=127.0.0.1 базовое знание
	DB_USER=
	DB_PASSWORD=
	DB_NAME=
	DB_PORT=5432 базовое знание
	```
4. Запустить `migrations.sql`
5. Запустить `go run main.go` или `./main`

# API

Можно импортировать запросы из `TodoPS.postman_collection.json`

Список запросов:

1. Показать общий список todos. `GET`: /todos

2. конкретный один todo по id. `GET`: /todo/:id

3. Сохранение todo. `POST`: /todo

4. Редактировать todo. `PUT`: /todo/:id

5. Удалить todo `DELETE`: /todo/:id

# Структура сущности
```go
type Todo struct {
	Id          int       `gorm:"primaryKey"`
	Description string    `json:"description"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
}
```


<details>
<summary><b>Как можно было бы улучшить проект</b></summary>

Заменить скрипт `migration.sql` на `Automigrate` от `gorm`

Добавить в файл `db.go`
```go
func (d *DB) Migrate() error {
	return d.db.AutoMigrate(
		&models.Todo{},
	)
}
```

Добавить в файл `main.go`, после database config
```go
if err = database.Migrate(); err != nil {
	log.Fatal(err)
}
```

</details>
