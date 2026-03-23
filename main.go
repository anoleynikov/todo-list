package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"todo-list/storage"
)

func main() {
	// Инициализация хранилища
	store, err := storage.NewStorage()
	if err != nil {
		fmt.Printf("Ошибка инициализации: %v\n", err)
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("=== Todo List ===")
	fmt.Println("Доступные команды:")
	fmt.Println("  add <текст>     - добавить задачу")
	fmt.Println("  list            - показать все задачи")
	fmt.Println("  done <id>       - отметить задачу как выполненную")
	fmt.Println("  toggle <id>     - переключить статус задачи")
	fmt.Println("  delete <id>     - удалить задачу")
	fmt.Println("  exit            - выход")
	fmt.Println()

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка ввода:", err)
			continue
		}

		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		parts := strings.Fields(input)
		command := parts[0]

		switch command {
		case "exit":
			fmt.Println("До свидания!")
			return

		case "add":
			if len(parts) < 2 {
				fmt.Println("Ошибка: укажите текст задачи")
				continue
			}
			title := strings.Join(parts[1:], " ")
			if err := store.Add(title); err != nil {
				fmt.Printf("Ошибка добавления: %v\n", err)
			} else {
				fmt.Println("✓ Задача добавлена")
			}

		case "list":
			tasks := store.List()
			if len(tasks) == 0 {
				fmt.Println("Список задач пуст")
			} else {
				fmt.Println("\nВаши задачи:")
				for _, t := range tasks {
					fmt.Println("  " + t.String())
				}
				fmt.Println()
			}

		case "done":
			if len(parts) < 2 {
				fmt.Println("Ошибка: укажите ID задачи")
				continue
			}
			id, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("Ошибка: ID должен быть числом")
				continue
			}
			if err := store.Complete(id); err != nil {
				fmt.Printf("Ошибка: %v\n", err)
			} else {
				fmt.Printf("✓ Задача %d отмечена как выполненная\n", id)
			}

		case "toggle":
			if len(parts) < 2 {
				fmt.Println("Ошибка: укажите ID задачи")
				continue
			}
			id, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("Ошибка: ID должен быть числом")
				continue
			}
			if err := store.Toggle(id); err != nil {
				fmt.Printf("Ошибка: %v\n", err)
			} else {
				fmt.Printf("✓ Статус задачи %d изменен\n", id)
			}

		case "delete":
			if len(parts) < 2 {
				fmt.Println("Ошибка: укажите ID задачи")
				continue
			}
			id, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("Ошибка: ID должен быть числом")
				continue
			}
			if err := store.Delete(id); err != nil {
				fmt.Printf("Ошибка: %v\n", err)
			} else {
				fmt.Printf("✓ Задача %d удалена\n", id)
			}

		default:
			fmt.Printf("Неизвестная команда: %s\n", command)
			fmt.Println("Доступные команды: add, list, done, toggle, delete, exit")
		}
	}
}
