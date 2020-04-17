package models

import "database/sql"

func MakeDbPerson() *DbPerson {
	return &DbPerson{}
}

type (
	Person struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}

	DbPerson struct {
		Id   *sql.NullInt32  `json:"id"`
		Name *sql.NullString `json:"name"`
	}
)

func GetPersonsTobeDeletedAndAdded(existingPerson, newPerson []Person) (personsToBeUpdated, personsToBeAdded, personsToBeDeleted []Person) {

	m := make(map[int]bool)
	idPersonMap := make(map[int]Person)

	for _, item := range existingPerson {
		m[item.Id] = true
		idPersonMap[item.Id] = item
	}
	for _, item := range newPerson {
		if item.Id == 0 {
			personsToBeAdded = append(personsToBeAdded, item)
		} else if _, ok := m[item.Id]; ok {
			personsToBeUpdated = append(personsToBeUpdated, item)
			delete(m, item.Id)
		} else {
			personsToBeDeleted = append(personsToBeDeleted, item)
		}

	}

	for k, _ := range m {
		personsToBeDeleted = append(personsToBeDeleted, idPersonMap[k])
	}
	return

}

func FromDbPerson(dbPerson *DbPerson) (person *Person) {
	person = &Person{}
	if dbPerson.Id != nil && dbPerson.Id.Valid {
		person.Id = int(dbPerson.Id.Int32)
	}
	if dbPerson.Name != nil && dbPerson.Name.Valid {
		person.Name = dbPerson.Name.String
	}
	return
}

func ToDbPerson(person *Person) (personDb *DbPerson) {
	personDb = &DbPerson{}
	personDb.Id = getNullInt(person.Id)
	personDb.Name = getNullString(person.Name)
	return
}
