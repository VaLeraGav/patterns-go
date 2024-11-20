package future_promise

type Result struct {
	value int
	err   error
}

func Promise(task func() (int, error)) chan Result {
	resultCh := make(chan Result, 1) // создаем канал для результата

	go func() {
		value, err := task()                       // выполняем задачу
		resultCh <- Result{value: value, err: err} // отправляем результат и ошибку в канал
		close(resultCh)                            // закрываем канал
	}()

	return resultCh
}
