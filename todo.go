package main

// Todo is a task to do.
type Todo struct {
	Description string // Content of the todo
	Date        string // Timestamp from time of creation
	Done        bool   // Complete or not
}

// Toggle the `Done` attribute of a `Todo`.
func (t *Todo) toggle() {
	t.Done = !t.Done
}

// Update the `Description` attribute of a `Todo`.
func (t *Todo) update(description string) {
	t.Description = description
}

// Todos is a slice of `todo`
type Todos []Todo

// Toggle the `Done` attribute of a Todo.
func Toggle(date string, tds Todos) {
	index := getIndex(date, tds)
	td := tds[index]
	td.toggle()
}

// Drop a todo from Todos.
// This is a poor way to go about this. Needs much improvement.
func Drop(date string, tds Todos) Todos {
	var result Todos
	for _, v := range tds {
		if v.Date != date {
			result = append(result, v)
		}
	}
	return result
}

// Drop a todo from Todos.
// func Drop(date string, tds Todos) Todos {
// 	index := getIndex(date, tds)
// 	below := index - 1
// 	above := index + 1
// 	return append(tds[:below], tds[above:]...)
// }

func getIndex(date string, todos Todos) int {
	var res int
	for i, todo := range todos {
		if todo.Date == date {
			res = i
		}
	}
	return res
}
