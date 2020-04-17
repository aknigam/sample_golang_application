package models

import "database/sql"
import "time"
import "github.com/go-sql-driver/mysql"

func MakeDbSprint() *DbSprint {
	return &DbSprint{
		Owner: &DbPerson{
			Id: &sql.NullInt32{},
		},
	}
}

type (
	Sprint struct {
		Id               int        `json:"id"`
		Title            string     `json:"title"`
		Stories          []Story    `json:"stories"`
		Owner            *Person    `json:"owner"`
		StartDate        *time.Time `json:"startDate"`
		StartTime        string     `json:"startTime"`
		EndDateTimestamp *time.Time `json:"endDateTimestamp"`
		CurrentDatetime  *time.Time `json:"currentDatetime"`
		IsActive         bool       `json:"isActive"`
	}

	DbSprint struct {
		Id               *sql.NullInt32  `json:"id"`
		Title            *sql.NullString `json:"title"`
		Stories          []DbStory       `json:"stories"`
		Owner            *DbPerson       `json:"owner"`
		StartDate        *mysql.NullTime `json:"startDate"`
		StartTime        string          `json:"startTime"`
		EndDateTimestamp *mysql.NullTime `json:"endDateTimestamp"`
		CurrentDatetime  *mysql.NullTime `json:"currentDatetime"`
		IsActive         *sql.NullBool   `json:"isActive"`
	}
)

func GetSprintsTobeDeletedAndAdded(existingSprint, newSprint []Sprint) (sprintsToBeUpdated, sprintsToBeAdded, sprintsToBeDeleted []Sprint) {

	m := make(map[int]bool)
	idSprintMap := make(map[int]Sprint)

	for _, item := range existingSprint {
		m[item.Id] = true
		idSprintMap[item.Id] = item
	}
	for _, item := range newSprint {
		if item.Id == 0 {
			sprintsToBeAdded = append(sprintsToBeAdded, item)
		} else if _, ok := m[item.Id]; ok {
			sprintsToBeUpdated = append(sprintsToBeUpdated, item)
			delete(m, item.Id)
		} else {
			sprintsToBeDeleted = append(sprintsToBeDeleted, item)
		}

	}

	for k, _ := range m {
		sprintsToBeDeleted = append(sprintsToBeDeleted, idSprintMap[k])
	}
	return

}

func FromDbSprint(dbSprint *DbSprint) (sprint *Sprint) {
	sprint = &Sprint{}
	if dbSprint.Owner != nil && dbSprint.Owner.Id != nil && dbSprint.Owner.Id.Valid {
		sprint.Owner = &Person{
			Id: int(dbSprint.Owner.Id.Int32),
		}
	}
	if dbSprint.Id != nil && dbSprint.Id.Valid {
		sprint.Id = int(dbSprint.Id.Int32)
	}
	if dbSprint.Title != nil && dbSprint.Title.Valid {
		sprint.Title = dbSprint.Title.String
	}
	if dbSprint.StartDate != nil && dbSprint.StartDate.Valid {
		sprint.StartDate = &dbSprint.StartDate.Time
	}
	sprint.StartTime = dbSprint.StartTime
	if dbSprint.EndDateTimestamp != nil && dbSprint.EndDateTimestamp.Valid {
		sprint.EndDateTimestamp = &dbSprint.EndDateTimestamp.Time
	}
	if dbSprint.CurrentDatetime != nil && dbSprint.CurrentDatetime.Valid {
		sprint.CurrentDatetime = &dbSprint.CurrentDatetime.Time
	}
	if dbSprint.IsActive != nil && dbSprint.IsActive.Valid {
		sprint.IsActive = dbSprint.IsActive.Bool
	}
	return
}

func ToDbSprint(sprint *Sprint) (sprintDb *DbSprint) {
	sprintDb = &DbSprint{}
	sprintDb.Id = getNullInt(sprint.Id)
	sprintDb.Title = getNullString(sprint.Title)
	sprintDb.StartDate = getNullTime(sprint.StartDate)
	sprintDb.StartTime = sprint.StartTime
	sprintDb.EndDateTimestamp = getNullTime(sprint.EndDateTimestamp)
	sprintDb.CurrentDatetime = getNullTime(sprint.CurrentDatetime)
	sprintDb.IsActive = getNullBool(sprint.IsActive)
	return
}
