package models

type sessionForeignKey struct{}

/*

	Add the session's user_id forein key

*/
func (s sessionForeignKey) up(db *DB) error {

	db.Client.Model(&Session{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	return nil

}

/*

	Name of this migration

*/
func (s sessionForeignKey) getName() string {

	return "$sessionForeignKey"

}
