package singleton

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single

func GetInstanceMutex() *single {

	// В начале нужна nil-проверка, с ее помощью мы убеждаемся, что первый экземпляр singleInstance — пустой.
	// Благодаря этому мы можем избежать ресурсоемких операций блокировки при каждом вызове getInstance.'
	// Если эта проверка не пройдена, тогда поле singleInstance уже заполнено.
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		// После блокировки используется еще одна nil-проверка.
		// В случаях, когда первую проверку проходит более одного потока, вторая обеспечивает создание экземпляра одиночки единым потоком.
		// В противном случае, все потоки создавали бы свои экземпляры структуры одиночки.
		if singleInstance == nil {
			fmt.Println("Creating single instance now.")

			// Структура singleInstance создается внутри блокировки.
			singleInstance = &single{}
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}
