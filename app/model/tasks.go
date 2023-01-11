package model

// Tasks structure for our blog
type Tasks struct {
	ID       uint64 `json:"id"`
	Text     string `json:"text"`
	Day      string `json:"day"`
	Reminder bool   `json:"reminder"`
}

func GetAllTasks() ([]Tasks, error) {
	var tasks []Tasks

	query := `select id, text, day, reminder from tasks;`

	rows, err := db.Query(query)
	if err != nil {
		return tasks, err
	}

	defer rows.Close()

	for rows.Next() {
		var id uint64
		var text, day string
		var reminder bool

		err := rows.Scan(&id, &text, &day, &reminder)
		if err != nil {
			return tasks, err
		}

		task := Tasks{
			ID:       id,
			Text:     text,
			Day:      day,
			Reminder: reminder,
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func GetTask(id uint64) (Tasks, error) {
	var task Tasks

	query := `select text, day, reminder from tasks where id=$1`
	row, err := db.Query(query, id)
	if err != nil {
		return task, err
	}

	defer row.Close()

	if row.Next() {
		var text, day string
		var reminder bool

		err := row.Scan(&text, &day, &reminder)
		if err != nil {
			return task, err
		}

		task = Tasks{
			ID:       id,
			Text:     text,
			Day:      day,
			Reminder: reminder,
		}
	}

	return task, nil
}

func CreateTask(task Tasks) (Tasks, error) {

	query := `insert into tasks(text, day,reminder) values($1, $2,$3) returning id;`

	//d_, err := db.Exec(query, task.Text, task.Day, task.Reminder)

	var id uint64
	err := db.QueryRow(query, task.Text, task.Day, task.Reminder).Scan(&id)

	if err != nil {
		return task, err
	}
	task.ID = id

	return task, nil
}

func DeleteTask(id uint64) error {
	query := `delete from tasks where id=$1;`
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func UpdateTask(task Tasks) error {

	query := `update tasks set text=$1, day=$2, reminder=$3 where id=$4;`

	_, err := db.Exec(query, task.Text, task.Day, task.Reminder, task.ID)
	if err != nil {
		return err
	}
	return nil
}
