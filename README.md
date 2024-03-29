# CLI To-Do List App

This is a simple command-line interface (CLI) to-do list application written in Golang. This project's purpose is to demonstrate implementing CRUD functionality into Golang CLI applications. The application allows users to manage their to-do tasks with basic CRUD (Create, Read, Update, Delete) operations.


## Features

- Add new tasks to the to-do list.
- View the list of tasks with corresponding numbers.
- Mark tasks as done.
- Delete tasks from the list.

## Getting Started

### Prerequisites

1. Clone the repository:
   ```bash
   git clone https://github.com/amy324/Golang-Todo-List-CLI-App.git
   ```
2.  Make sure you have [Go](https://golang.org/) installed on your machine.
3. Run the application using the terminal command
    ```
     go run main.go
    ```
    or alternatively:
    ```
     go run .
    ```
4. Follow the on-screen instructions to interact with the to-do list. 

### Using the app
Here's a quick overview of available options:

1. **Add Task:** Add a new task to the to-do list.
2. **View Tasks:** Display the list of tasks with corresponding numbers.
3. **Mark Task as Done:** Mark a specific task as done by entering its number.
4. **Delete Task:** Delete a specific task by entering its numbe.
5. **Exit:** Exit the application.

## Code Explanation

- **`main` function:** The entry point of the application that contains the main loop for user interaction.
- **`getUserInput` function:** Helper function to get user input from the command line.
- **CRUD operations:**
  - `addTask`: Adds a new task to the to-do list file.
  - `viewTasks`: Displays the tasks with corresponding numbers.
  - `markTaskAsDone`: Marks a task as done by modifying the task in the list.
  - `deleteTask`: Deletes a task from the list.

## Terminal Output Walkthrough

1. Expected output after running `go main.go`

```
Welcome to your CLI To-Do List

Options:
1. Add Task
2. View Tasks
3. Mark Task as Done
4. Delete Task
5. Exit
Select an option (1/2/3/4/5): 

```

2. Ouput after entering `1` to select the Add Task Option:

```
Options:
1. Add Task
2. View Tasks
3. Mark Task as Done
4. Delete Task
5. Exit
Select an option (1/2/3/4/5): 
1
Enter task: 
Do grocery shopping
Task added successfully!
```
3. Output after selecting option `2` to View Tasks
```Options:
1. Add Task
2. View Tasks
3. Mark Task as Done
4. Delete Task
5. Exit
Select an option (1/2/3/4/5): 
2

Tasks:
1. Do grocery shopping
2. Feed the cat
3. Water the plants

```


4. Output after marking a task as done by selecting option `3` and the option `2` to view tasks - task has now been marked as [DONE].

```Options:
Options:
1. Add Task
2. View Tasks
3. Mark Task as Done
4. Delete Task
5. Exit
Select an option (1/2/3/4/5): 
3
Enter the number of the task to mark as done: 
2

Task marked as done successfully!
1. Add Task
2. View Tasks
3. Mark Task as Done
4. Delete Task
5. Exit
Select an option (1/2/3/4/5): 
2

Tasks:
1. Do grocery shopping
2. [Done] Feed the cat
3. Water the plants


```
5. Output after selecting option `4` to delete a task and then option `2` - task has now been deleted
```Options:
1. Add Task
2. View Tasks
3. Mark Task as Done
4. Delete Task
5. Exit
Select an option (1/2/3/4/5): 
4
Enter the number of the task to delete: 
2
Task deleted successfully!

Options:
1. Add Task
2. View Tasks
3. Mark Task as Done
4. Delete Task
5. Exit
Select an option (1/2/3/4/5): 
2

Tasks:
1. Do grocery shopping
2. Water the plants

```
6. Output after selecting option `5` to exit the app
```
Options:
1. Add Task
2. View Tasks
3. Mark Task as Done
4. Delete Task
5. Exit
Select an option (1/2/3/4/5): 
5
Exiting CLI To-Do List. Goodbye!`
```
## Contributing

Contributions are welcome! If you have any ideas, suggestions, or improvements, feel free to open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).
