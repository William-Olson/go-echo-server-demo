package models

type addTestData struct{}

/*

	Add some test data to db

*/
func (m addTestData) up(db *DB) error {

	db.Users.Create("admin", "testFirst", "testLast", "testPassword")
	return nil

}

/*

	Name of this migration

*/
func (m addTestData) getName() string {

	return "$addTestData"

}
