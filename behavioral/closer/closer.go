package closer

import (
	"fmt"
	"strings"
	"sync"
)

type Closer struct {
	funcs []func() error
	wg    sync.WaitGroup
}

var (
	instance *Closer
	once     sync.Once
)

func GetInstance() *Closer {
	once.Do(func() {
		instance = &Closer{
			funcs: make([]func() error, 0),
		}
	})
	return instance
}

func (c *Closer) Add(f ...func() error) {
	c.funcs = append(c.funcs, f...)
}

func (c *Closer) CloseAll() error {
	var msgs = make([]string, 0, len(c.funcs))

	for _, f := range c.funcs {
		if err := f(); err != nil {
			msgs = append(msgs, fmt.Sprintf("[!] %v", err))
		}
	}

	if len(msgs) > 0 {
		return fmt.Errorf(
			"shutdown finished with error(s): \n%s",
			strings.Join(msgs, "\n"),
		)
	}

	return nil
}

func (c *Closer) CloseWg() error {
	var msgs = make([]string, 0, len(c.funcs))

	for _, f := range c.funcs {
		c.wg.Add(1) // Увеличиваем счетчик WaitGroup
		go func(fn func() error) {
			defer c.wg.Done() // Уменьшаем счетчик при завершении горутины
			if err := fn(); err != nil {
				msgs = append(msgs, fmt.Sprintf("[!] %v", err))
			}
		}(f) // Передаем f как аргумент
	}

	c.wg.Wait() // Ожидаем завершения всех горутин

	if len(msgs) > 0 {
		return fmt.Errorf(
			"shutdown finished with error(s): \n%s",
			strings.Join(msgs, "\n"),
		)
	}

	return nil
}
