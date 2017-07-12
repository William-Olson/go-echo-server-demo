package models

type addTestData struct{}

func (m addTestData) up(db *DB) error {

	// add test data to db

	db.Users.Create("admin", "testFirst", "testLast", "testPassword")
	return nil

}

func (m addTestData) getName() string {

	return "$addTestData"

}
