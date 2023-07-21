package database

type Todo struct {
	Id    string
	Title string
	Done  bool
}

func GetTodos() ([]Todo, error) {
	db, err := setupDatabase()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, todo, completed FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var arrTodos []Todo
	for rows.Next() {
		var id string
		var todo string
		var completed bool
		err := rows.Scan(&id, &todo, &completed)
		if err != nil {
			return nil, err
		}
		arrTodos = append(arrTodos, Todo{
			Id:    id,
			Title: todo,
			Done:  completed,
		})
	}
	return arrTodos, nil
}

func SaveTodo(todo string, completed bool) (string, error) {
	db, err := setupDatabase()
	if err != nil {
		return "", err
	}
	defer db.Close()
	var id string
	err = db.QueryRow("INSERT INTO todos (todo, completed) VALUES ($1, $2) RETURNING id", todo, completed).Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func UpdateTodo(id string, todo string, completed bool) error {
	db, err := setupDatabase()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("UPDATE todos SET todo = $1, completed = $2 WHERE id = $3", todo, completed, id)
	if err != nil {
		return err
	}

	return nil
}

func DeleteTodo(id string) error {
	db, err := setupDatabase()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM todos WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
