package repository

import (
	_ "github.com/go-sql-driver/mysql"
	"sample_golang_application/common"
	"sample_golang_application/models"
)

type TaskRepository struct {
}

func (repo *TaskRepository) GetTask(taskId int) (task *models.Task, err error) {
	dbtask := models.MakeDbTask()
	txn, err := Db.Begin()
	if err != nil {
		return
	}
	defer txn.Commit()
	err = Db.QueryRow(" SELECT id, name, story_id"+
		" FROM task"+
		" WHERE (id = ?)", taskId).
		Scan(&dbtask.Id, &dbtask.Name, &dbtask.StoryId)
	if err != nil {
		txn.Rollback()
	}
	return models.FromDbTask(dbtask), err
}

func (repo *TaskRepository) DeleteTask(taskId int) (err error) {
	_, err = Db.Exec(" DELETE FROM task"+
		" WHERE (id =  ?)", taskId)
	return
}

func (repo *TaskRepository) UpdateTask(task *models.Task) (err error) {
	dbtask := models.ToDbTask(task)
	_, err = Db.Exec(" UPDATE task"+
		" SET name = ?, story_id = ?"+
		" WHERE (id = ?)", dbtask.Name, dbtask.StoryId, dbtask.Id)
	if err != nil {
		common.Error.Println("Task could not be updated ")
		return
	}
	return
}

func (repo *TaskRepository) CreateTask(task *models.Task) (id int, err error) {
	dbtask := models.ToDbTask(task)
	statement := " INSERT INTO task" +
		"  (name, story_id)" +
		" VALUES (?, ?)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(dbtask.Name, dbtask.StoryId)

	if err != nil {
		common.Error.Println("Could not create Task ", err)
		return
	}
	generatedId, err := result.LastInsertId()
	if err != nil {
		common.Error.Println("Could not create Task ", err)
		return
	}
	return int(generatedId), err

}

func (repo *TaskRepository) GetAllStoryTasks(storyId int) (tasks []models.Task) {
	rows, err := Db.Query(" SELECT id, name, story_id"+
		" FROM task"+
		" WHERE (story_id = ?)", storyId)
	if err != nil {
		common.Error.Println("Could not find tasks ", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		task := models.Task{}
		err := rows.Scan(&task.Id, &task.Name, &task.StoryId)
		if err != nil {
			common.Error.Println("Could not find tasks ", err)
			break
		}
		tasks = append(tasks, task)
	}
	err = rows.Err()
	if err != nil {
		common.Error.Fatal(err)
	}
	return
}
