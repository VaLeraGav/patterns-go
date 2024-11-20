# Паттерны многопоточности (Parallel)

`Параллелизм` — это основная функция языка программирования Go, и понимание того, как эффективно его использовать, необходимо для создания эффективных и масштабируемых систем.

- [Fan-In Fan-Out](faninandfanout)
- [Обещание (Future/Promise)](future_promise)
- [Генератор (Generator)](generator)
- [Map and Filter](map_filter)
- [Pipeline](pipeline)
- [Семафор (Semaphore)](semaphore)
- [Подписка (Subscription)](subscription)
- [Пула рабочих (Workerpool)](workerpool)

> Применение многопоточности без необходимости может не только усложнить код, но и сделать его менее читаемым, замедлить выполнение и привести к ошибкам. Слепое использование паттернов, там где можно обойтись простым синхронным выполнением, часто создает больше проблем, чем решает.

## Работа с Ошибками в многопоточности

Рассмотрим как эффективно передавать как результаты, так и ошибки между горутинами, обеспечивая чистоту и читаемость кода.

Теперь мы будем передавать ошибки из горутин через канал с исполдьзованием структуры. Эта структура используется для хранения результата (`data`) и ошибки (`err`). Это позволяет удобно передавать как данные, так и информацию об ошибках через канал.

```go
// Result — структура для хранения результата и ошибки
type Result struct {
    data int
    err  error
}

func main() {
    input := []int{1, 2, 3, 4}

    resultCh := make(chan Result)

    // запускаем потребителя, который будет отправлять результаты и ошибки
    go consumer(generator(input), resultCh)

    // читаем результаты
    for res := range resultCh {
        if res.err != nil {
            log.Println("Ошибка:", res.err)
        } else {
            log.Println("Результат:", res.data)
        }
    }
}

// generator отправляет данные в канал
func generator(input []int) chan int {
    inputCh := make(chan int)

    go func() {
        defer close(inputCh)
        for _, data := range input {
            inputCh <- data
        }
    }()

    return inputCh
}

// consumer вызывает функцию, которая может возвращать ошибку
func consumer(inputCh chan int, resultCh chan Result) {
    defer close(resultCh)

    for data := range inputCh {
        resp, err := callDatabase(data)
        resultCh <- Result{data: resp, err: err}
    }
}

// callDatabase возвращает ошибку
func callDatabase(data int) (int, error) {
    return data, errors.New("ошибка запроса к базе данных")
}
```

Однако, есть еще `errgroup` из пакета `golang.org/x/sync/errgroup` который предоставляет синхронизацию, распространение ошибок и отмену контекста для групп горутин.

## Конструкция For-Select-Done

Поскольку мы оперируем несколькими горутинами, нам нужно быть особенно внимательными, чтобы не допустить никаких утечек памяти в нашей программе. На всякий случай напомню, что утечка происходит всякий раз, когда горутина остается висеть в нашей программе после того, как она нам больше не нужна. Решение заключается в том, чтобы посылать горутине сигнал, который дает ей знать, что она может завершиться.

Наиболее распространенным способом реализации этого механизма является объединение цикла for-select с отдельным каналом, который отправляет сигнал завершения горутине. Обычно его называют done-каналом.

```go
func printIntegers(done <-chan struct{}, intStream <-chan int) {
  for{
    select {
      case i := <-intStream:
        fmt.Println(i)
      case <-done:
        return
    }
  }
}
```

Следует отметить, что мы можем добиться того же эффекта, используя пакет Context:

```go
func printIntegers(ctx context.Context, intStream <-chan int) {
  for{
    select {
      case i := <-intStream:
        fmt.Println(i)
      case <-ctx.Done():
        return
    }
  }
}
```
