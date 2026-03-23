package task

import (
	"fmt"
	"time"
)

// Task представляет собой одну задачу
type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

// NewTask создает новую задачу
func NewTask(id int, title string) *Task {
	return &Task{
		ID:        id,
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
	}
}

// String возвращает строковое представление задачи
func (t *Task) String() string {
	status := " "
	if t.Completed {
		status = "✓"
	}
	return fmt.Sprintf("[%s] %d. %s (создано: %s)",
		status, t.ID, t.Title, t.CreatedAt.Format("02.01.2006 15:04"))
}

// Toggle переключает статус выполнения задачи
func (t *Task) Toggle() {
	t.Completed = !t.Completed
}
