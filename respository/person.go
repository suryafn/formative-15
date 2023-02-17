package respository

import (
	"database/sql"
	"formative-15/structs"
)

func GetAllPerson(db *sql.DB) (err error, results []structs.Person) {
	sql := "select * from person"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var person = structs.Person{}

		err = rows.Scan(&person.ID, &person.FirstName, &person.LastName)

		if err != nil {
			panic(err)
		}

		results = append(results, person)
	}

	return
}

func InsertPerson(db *sql.DB, person structs.Person) (err error) {
	sql := "insert into person (first_name, last_name) values ($1, $2)"

	errs := db.QueryRow(sql, person.FirstName, person.LastName)

	return errs.Err()
}

func UpdatePerson(db *sql.DB, person structs.Person) (err error) {
	sql := "update person set first_name = $2,last_name= $3 where id = $1 "
	errs := db.QueryRow(sql, person.ID, person.FirstName, person.LastName)

	return errs.Err()
}

func DeletePerson(db *sql.DB, person structs.Person) (err error) {
	sql := "delete from person where id= $1"
	errs := db.QueryRow(sql, person.ID)

	return errs.Err()
}
