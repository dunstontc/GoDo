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
func Toggle(tds Todos, date string) {
	index := getIndex(tds, date)
	td := tds[index]
	td.toggle()
}

// // Update the `Description` attribute of a Todo.
// func Update(tds Todos, date string, desc string) {
// 	index := getIndex(tds, date)
// 	td := tds[index]
// 	td.update(desc)
// }

// Drop a todo from Todos.
// This is a poor way to go about this. Needs much improvement.
func Drop(tds Todos, date string) Todos {
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

// getIndex returns a `todo`'s position in a `Todos` slice.
func getIndex(todos Todos, date string) int {
	var res int
	for i, todo := range todos {
		if todo.Date == date {
			res = i
		}
	}
	return res
}
