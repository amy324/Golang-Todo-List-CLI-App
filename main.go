package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/fatih/color"
)

func main() {
	// Print welcome message
	color.Cyan("Welcome to your CLI To-Do List")

	// Infinite loop for user interaction
	for {
		color.Cyan("\nOptions:")
		fmt.Println("1. Add Task")
		fmt.Println("2. View Tasks")
		fmt.Println("3. Mark Task as Done")
		fmt.Println("4. Delete Task")
		fmt.Println("5. Exit")

		color.Cyan("Select an option (1/2/3/4/5): ")
		option := getUserInput()

		switch option {
		case "1":
			addTask()
		case "2":
			viewTasks()
		case "3":
			markTaskAsDone()
		case "4":
			deleteTask()
		case "5":
			color.Cyan("Exiting CLI To-Do List. Goodbye!")
			os.Exit(0)
		default:
			color.Red("Invalid option. Please try again.")
		}
	}
}
// Function to get user input
func getUserInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
//Function to add a tasks
func addTask() {
	color.Cyan("Enter task: ")
	task := getUserInput()

	file, err := os.OpenFile("tasks.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		color.Red("Error opening file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(task + "\n")
	if err != nil {
		color.Red("Error writing to file:", err)
		return
	}

	color.Green("Task added successfully!")
}

// Function to view tasks in the to-do list
func viewTasks() {
	file, err := os.Open("tasks.txt")
	if err != nil {
		color.Red("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	tasks := make([]string, 0)

	// Read and print tasks from the file
	color.Yellow("\nTasks:")
	for scanner.Scan() {
		task := scanner.Text()
		tasks = append(tasks, task)
	}

	if err := scanner.Err(); err != nil {
		color.Red("Error reading file:", err)
		return
	}

	for i, task := range tasks {
		color.Yellow("%d. %s\n", i+1, task)
	}
}

//Function to mark tasks as done
func markTaskAsDone() {
	color.Cyan("Enter the number of the task to mark as done: ")
	taskNumberStr := getUserInput()
	taskNumber, err := strconv.Atoi(taskNumberStr)
	if err != nil {
		color.Red("Invalid input. Please enter a valid task number.")
		return
	}

	tasks, err := readTasksFromFile()
	if err != nil || taskNumber < 1 || taskNumber > len(tasks) {
		fmt.Println("Invalid task number. Please try again.")
		return
	}

	// Mark the selected task as done
	tasks[taskNumber-1] = "[Done] " + tasks[taskNumber-1]

	// Write the updated tasks back to the file
	writeTasksToFile(tasks)

	color.Green("Task marked as done successfully!")
}

//Function to delete tasks
func deleteTask() {
	color.Cyan("Enter the number of the task to delete: ")
	taskNumberStr := getUserInput()
	taskNumber, err := strconv.Atoi(taskNumberStr)
	if err != nil {
		color.Red("Invalid input. Please enter a valid task number.")
		return
	}

	tasks, err := readTasksFromFile()
	if err != nil || taskNumber < 1 || taskNumber > len(tasks) {
		color.Red("Invalid task number. Please try again.")
		return
	}

	// Remove the selected task
	tasks = append(tasks[:taskNumber-1], tasks[taskNumber:]...)

	// Write the updated tasks back to the file
	writeTasksToFile(tasks)

	color.Green("Task deleted successfully!")
}

//Function to read the tasks from .txt file
func readTasksFromFile() ([]string, error) {
	file, err := os.Open("tasks.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var tasks []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tasks = append(tasks, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func writeTasksToFile(tasks []string) error {
	file, err := os.Create("tasks.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	for _, task := range tasks {
		_, err := file.WriteString(task + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}
