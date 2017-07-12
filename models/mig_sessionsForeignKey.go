package models

type sessionForeignKey struct{}

func (s sessionForeignKey) up(db *DB) error {

	// add the session's user_id forein key
	db.Client.Model(&Session{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")

	return nil

}

func (s sessionForeignKey) getName() string {

	return "$sessionForeignKey"

}
