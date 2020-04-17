package repository

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"sample_golang_application/common"
	"sample_golang_application/models"
)

type SprintRepository struct {
}

func (repo *SprintRepository) GetSprint(sprintId int) (sprint *models.Sprint, err error) {
	dbsprint := models.MakeDbSprint()
	//txn, err := Db.Begin()
	//if err != nil {
	//	return
	//}
	//defer txn.Commit()
	err = Db.QueryRow(" SELECT id, title, owner, start_date, start_time, end_date_timestamp, current_datetime, is_active"+
		" FROM sprint"+
		" WHERE (id = ?)", sprintId).
		Scan(&dbsprint.Id, &dbsprint.Title, &dbsprint.Owner.Id, &dbsprint.StartDate, &dbsprint.StartTime, &dbsprint.EndDateTimestamp, &dbsprint.CurrentDatetime, &dbsprint.IsActive)
	//if err != nil {
	//	txn.Rollback()
	//}
	return models.FromDbSprint(dbsprint), err
}

func (repo *SprintRepository) DeleteSprint(sprintId int) (err error) {
	_, err = Db.Exec(" DELETE FROM sprint"+
		" WHERE (id =  ?)", sprintId)
	return
}

func (repo *SprintRepository) UpdateSprint(sprint *models.Sprint) (err error) {
	txn, err := Db.Begin()
	if err!= nil {
		return
	}
	defer txn.Commit()
	dbsprint := models.ToDbSprint(sprint)
	result, err := txn.Exec(" UPDATE sprint"+
		" SET title = ?, start_date = ?, start_time = ?, end_date_timestamp = ?, current_datetime = ?, is_active = ?"+
		" WHERE (id = ?)", dbsprint.Title, dbsprint.StartDate, dbsprint.StartTime, dbsprint.EndDateTimestamp, dbsprint.CurrentDatetime, dbsprint.IsActive, dbsprint.Id)
	if err != nil {
		common.Error.Println("Sprint could not be updated ")
		txn.Rollback()
		return
	}

	fmt.Println(result.RowsAffected())
	return
}

func (repo *SprintRepository) CreateSprint(sprint *models.Sprint) (id int, err error) {
	dbsprint := models.ToDbSprint(sprint)
	statement := " INSERT INTO sprint" +
		"  (title, start_date, start_time, end_date_timestamp, current_datetime, is_active)" +
		" VALUES (?, ?, ?, ?, ?, ?)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(dbsprint.Title, dbsprint.StartDate, dbsprint.StartTime, dbsprint.EndDateTimestamp, dbsprint.CurrentDatetime, dbsprint.IsActive)

	if err != nil {
		common.Error.Println("Could not create Sprint ", err)
		return
	}
	generatedId, err := result.LastInsertId()
	if err != nil {
		common.Error.Println("Could not create Sprint ", err)
		return
	}
	return int(generatedId), err

}

func (repo *SprintRepository) LinkSprintOwner(sprintId int, ownerId int) (err error) {
	_, err = Db.Exec(" UPDATE sprint"+
		" SET owner = ?"+
		" WHERE (id = ? )", ownerId, sprintId)
	if err != nil {
		common.Error.Println("Failed to link owner from sprint ")
		panic(err)
	}
	return
}
func (repo *SprintRepository) UnlinkSprintOwner(sprintId int, ownerId int) (err error) {
	_, err = Db.Exec(" UPDATE sprint"+
		" SET owner = null"+
		" WHERE (id= ?)", sprintId)
	if err != nil {
		common.Error.Println("Failed to unlink owner from sprint ")
		panic(err)
	}
	return
}
func (repo *SprintRepository) LinkSprintStories(sprintId int, storieId int) (err error) {
	_, err = Db.Exec(" UPDATE story"+
		" SET sprint_id= ?"+
		" WHERE (id= ?)", sprintId, storieId)
	if err != nil {
		common.Error.Println("Failed to link storie with sprint ")
		panic(err)
	}
	return
}
func (repo *SprintRepository) UnlinkSprintStories(sprintId int, storieId int) (err error) {
	_, err = Db.Exec(" UPDATE story"+
		" SET sprint_id= null"+
		" WHERE (id= ?)", storieId)
	if err != nil {
		common.Error.Println("Failed to unlink stories from sprint ")
		panic(err)
	}
	return
}
func (repo *SprintRepository) UnlinkAllSprintStories(sprintId int) (err error) {
	_, err = Db.Exec(" UPDATE story"+
		" SET sprint_id= null"+
		" WHERE (sprint_id= ?)", sprintId)
	if err != nil {
		common.Error.Println("Could not delete stories ")
		panic(err)
	}
	return
}
func (repo *SprintRepository) DeleteAllSprintStories(sprintId int) (err error) {
	_, err = Db.Exec("delete from story where sprint_id = ?", sprintId)
	if err != nil {
		common.Error.Println("stories could not be deleted ")
		panic(err)
	}
	return
}
