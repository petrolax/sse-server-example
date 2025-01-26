package sse

import "fmt"

type Event struct {
	Event string // наименование события // если имя события не указано, то по умолчанию оно является message
	Data  []byte // данные события
	ID    string // Идентификатор события. Сделан строкой чтобы можно было использовать uuid
	Retry int    // период ожидания перед повторным подклчением в миллисекундах
}

// Каждое поле сообщения SSE разделяется символом \n, а конец сообщения означает \n\n
func (e *Event) String() string {
	result := ""
	if len(e.Event) != 0 {
		result += fmt.Sprintf("event: %s\n", e.Event)
	}
	if len(e.Data) != 0 {
		result += fmt.Sprintf("data: %s\n", e.Data)
	}
	if len(e.ID) != 0 {
		result += fmt.Sprintf("id: %s\n", e.ID)
	}
	if e.Retry != 0 {
		result += fmt.Sprintf("retry: %d\n", e.Retry)
	}
	return result + "\n"
}

func OpenEvent() *Event {
	return &Event{
		Event: "open",
		Retry: RetryDuration,
	}
}

func CloseEvent() *Event {
	return &Event{
		Event: "close",
		Retry: RetryDuration,
	}
}

func ErrorEvent() *Event {
	return &Event{
		Event: "error",
		Retry: RetryDuration,
	}
}
