package semaphore

// Semaphore — структура для управления количеством параллельных горутин
type Semaphore struct {
	semaCh chan struct{}
}

// NewSemaphore — создает новый семафор с заданной максимальной емкостью
func NewSemaphore(maxReq int) *Semaphore {
	// Счетчик семафора определяет, сколько горутин могут одновременно работать с ресурсом.
	return &Semaphore{
		semaCh: make(chan struct{}, maxReq),
	}
}

// Acquire — резервирует место в семафоре
func (s *Semaphore) Acquire() {
	s.semaCh <- struct{}{}
}

// Release — освобождает место в семафоре
func (s *Semaphore) Release() {
	<-s.semaCh
}
