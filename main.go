package main

import (
	"GreatProject/Business"
	"GreatProject/Infrastructure/Factory"
	"GreatProject/Infrastructure/Logger"
	"fmt"
)

func main() {
	logger := Logger.NewLogger()
	logger.SetLevel(Logger.DEBUG)

	repositoryFactory := Factory.NewRepositoryFactory()

	repository := repositoryFactory.GetTaskRepository(Factory.Postgres)

	taskService := Business.NewTaskService(repository)

	_, err := taskService.CreateTask("Go Basics", "Learn Go")
	if err != nil {
		fmt.Printf("Ошибка создания задачи: %v\n", err)
	}

	task2, err := taskService.CreateTask("Great Lesson", "First Great Lesson")
	if err != nil {
		fmt.Printf("Ошибка создания задачи: %v\n", err)
	}

	_, err = taskService.CreateTask("Next Lesson", "Next Great Lesson")
	if err != nil {
		fmt.Printf("Ошибка создания задачи: %v\n", err)
	}

	taskService.CompleteTask(task2.ID)

	fmt.Println("=== Все задачи ===")
	allTasks := taskService.GetAllTasks()
	for _, task := range allTasks {
		fmt.Printf("ID: %d, Name: %s, Description: %s, Completed: %t\n",
			task.ID, task.Name, task.Description, task.Completed)
	}

	fmt.Println("\n=== Выполненные задачи ===")
	completedTasks := taskService.GetCompletedTasks()
	for _, task := range completedTasks {
		fmt.Printf("ID: %d, Name: %s, Description: %s\n",
			task.ID, task.Name, task.Description)
	}

	fmt.Println("\n=== Ожидающие задачи ===")
	pendingTasks := taskService.GetPendingTasks()
	for _, task := range pendingTasks {
		fmt.Printf("ID: %d, Name: %s, Description: %s\n",
			task.ID, task.Name, task.Description)
	}
}
