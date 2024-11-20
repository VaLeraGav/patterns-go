package singleton

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/lib/pq" // Импортируем драйвер PostgreSQL
)

// dbSingleton - структура для хранения подключения к базе данных
type dbSingleton struct {
	connection *sql.DB
}

var instancedb *dbSingleton
var mu sync.Mutex

// GetInstance возвращает единственный экземпляр подключения к базе данных
func GetInstance(dataSourceName string) *dbSingleton {
	mu.Lock()
	defer mu.Unlock()

	if instance == nil {
		var err error
		instancedb = &dbSingleton{}
		instancedb.connection, err = sql.Open("postgres", dataSourceName)
		if err != nil {
			log.Fatalf("Ошибка подключения к базе данных: %v", err)
		}
	}

	return instancedb
}

// Метод для выполнения запроса
func (db *dbSingleton) Query(query string) (*sql.Rows, error) {
	return db.connection.Query(query)
}

// Метод для закрытия соединения
func (db *dbSingleton) Close() error {
	return db.connection.Close()
}

func Test() {
	dataSourceName := "user=username password=password dbname=mydb sslmode=disable"
	dbInstance := GetInstance(dataSourceName)

	// Пример выполнения запроса
	rows, err := dbInstance.Query("SELECT * FROM my_table")
	if err != nil {
		log.Fatalf("Ошибка выполнения запроса: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		// Обработка результата
	}

	// Закрытие соединения, если это необходимо
	err = dbInstance.Close()
	if err != nil {
		log.Fatalf("Ошибка закрытия соединения: %v", err)
	}
}
