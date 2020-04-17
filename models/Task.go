package models

import "database/sql"

func MakeDbTask() *DbTask {
	return &DbTask{}
}

type (
	Task struct {
		Id      int    `json:"id"`
		Name    string `json:"name"`
		StoryId int    `json:"storyId"`
	}

	DbTask struct {
		Id      *sql.NullInt32  `json:"id"`
		Name    *sql.NullString `json:"name"`
		StoryId *sql.NullInt32  `json:"storyId"`
	}
)

func GetTasksTobeDeletedAndAdded(existingTask, newTask []Task) (tasksToBeUpdated, tasksToBeAdded, tasksToBeDeleted []Task) {

	m := make(map[int]bool)
	idTaskMap := make(map[int]Task)

	for _, item := range existingTask {
		m[item.Id] = true
		idTaskMap[item.Id] = item
	}
	for _, item := range newTask {
		if item.Id == 0 {
			tasksToBeAdded = append(tasksToBeAdded, item)
		} else if _, ok := m[item.Id]; ok {
			tasksToBeUpdated = append(tasksToBeUpdated, item)
			delete(m, item.Id)
		} else {
			tasksToBeDeleted = append(tasksToBeDeleted, item)
		}

	}

	for k, _ := range m {
		tasksToBeDeleted = append(tasksToBeDeleted, idTaskMap[k])
	}
	return

}

func FromDbTask(dbTask *DbTask) (task *Task) {
	task = &Task{}
	if dbTask.Id != nil && dbTask.Id.Valid {
		task.Id = int(dbTask.Id.Int32)
	}
	if dbTask.Name != nil && dbTask.Name.Valid {
		task.Name = dbTask.Name.String
	}
	return
}

func ToDbTask(task *Task) (taskDb *DbTask) {
	taskDb = &DbTask{}
	taskDb.Id = getNullInt(task.Id)
	taskDb.Name = getNullString(task.Name)
	return
}
