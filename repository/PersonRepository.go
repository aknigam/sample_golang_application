package repository

import (
	_ "github.com/go-sql-driver/mysql"
	"sample_golang_application/common"
	"sample_golang_application/models"
)

type PersonRepository struct {
}

func (repo *PersonRepository) GetPerson(personId int) (person *models.Person, err error) {
	dbperson := models.MakeDbPerson()
	//txn, err := Db.Begin()
	//if err != nil {
	//	return
	//}
	//defer txn.Commit()
	err = Db.QueryRow(" SELECT id, name"+
		" FROM person"+
		" WHERE (id = ?)", personId).
		Scan(&dbperson.Id, &dbperson.Name)
	//if err != nil {
	//	txn.Rollback()
	//}
	return models.FromDbPerson(dbperson), err
}

func (repo *PersonRepository) DeletePerson(personId int) (err error) {
	_, err = Db.Exec(" DELETE FROM person"+
		" WHERE (id =  ?)", personId)
	return
}

func (repo *PersonRepository) UpdatePerson(person *models.Person) (err error) {
	dbperson := models.ToDbPerson(person)
	_, err = Db.Exec(" UPDATE person"+
		" SET name = ?"+
		" WHERE (id = ?)", dbperson.Name, dbperson.Id)
	if err != nil {
		common.Error.Println("Person could not be updated ")
		return
	}
	return
}

func (repo *PersonRepository) CreatePerson(person *models.Person) (id int, err error) {
	dbperson := models.ToDbPerson(person)
	statement := " INSERT INTO person" +
		"  (name)" +
		" VALUES (?)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(dbperson.Name)

	if err != nil {
		common.Error.Println("Could not create Person ", err)
		return
	}
	generatedId, err := result.LastInsertId()
	if err != nil {
		common.Error.Println("Could not create Person ", err)
		return
	}
	return int(generatedId), err

}
