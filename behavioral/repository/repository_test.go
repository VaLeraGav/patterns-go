package repository

import (
	"fmt"
	"testing"
)

func TestRepository(t *testing.T) {
	repo := NewInMemoryUserRepository()

	// Создаем нового пользователя
	user := &User{Name: "John Doe", Email: "john@example.com"}
	repo.Save(user)

	// Находим пользователя
	foundUser, err := repo.FindByID(user.ID)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Найден пользователь: %v\n", foundUser)

	// Удаляем пользователя
	err = repo.Delete(user.ID)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Пользователь удален.")

	// Проверяем, что пользователь удален
	_, err = repo.FindByID(user.ID)
	if err != nil {
		fmt.Println(err) // Ожидаем, что пользователь не найден
	}

	// 	Найден пользователь: &{1 John Doe john@example.com}
	// Пользователь удален.
	// user not found
}
