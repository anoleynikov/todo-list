package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"todo-list/task"
)

const (
	storageFile = "tasks.json"
)

// Storage управляет хранением задач
type Storage struct {
	tasks  []*task.Task
	nextID int
}

// NewStorage создает новое хранилище и загружает существующие задачи
func NewStorage() (*Storage, error) {
	s := &Storage{
		tasks:  make([]*task.Task, 0),
		nextID: 1,
	}

	if err := s.load(); err != nil {
		// Если файл не существует, это не ошибка
		if !os.IsNotExist(err) {
			return nil, err
		}
	}

	return s, nil
}

// load загружает задачи из файла
func (s *Storage) load() error {
	data, err := os.ReadFile(storageFile)
	if err != nil {
		return err
	}

	if len(data) == 0 {
		return nil
	}

	var tasks []*task.Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return err
	}

	s.tasks = tasks

	// Обновляем nextID
	maxID := 0
	for _, t := range s.tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	s.nextID = maxID + 1

	return nil
}

// save сохраняет задачи в файл
func (s *Storage) save() error {
	data, err := json.MarshalIndent(s.tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(storageFile, data, 0644)
}

// Add добавляет новую задачу
func (s *Storage) Add(title string) error {
	newTask := task.NewTask(s.nextID, title)
	s.tasks = append(s.tasks, newTask)
	s.nextID++

	return s.save()
}

// List возвращает все задачи
func (s *Storage) List() []*task.Task {
	return s.tasks
}

// Complete отмечает задачу как выполненную
func (s *Storage) Complete(id int) error {
	for _, t := range s.tasks {
		if t.ID == id {
			t.Completed = true
			return s.save()
		}
	}
	return fmt.Errorf("задача с ID %d не найдена", id)
}

// Toggle переключает статус задачи
func (s *Storage) Toggle(id int) error {
	for _, t := range s.tasks {
		if t.ID == id {
			t.Toggle()
			return s.save()
		}
	}
	return fmt.Errorf("задача с ID %d не найдена", id)
}

// Delete удаляет задачу
func (s *Storage) Delete(id int) error {
	for i, t := range s.tasks {
		if t.ID == id {
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			return s.save()
		}
	}
	return fmt.Errorf("задача с ID %d не найдена", id)
}
