package initialisation

import "database/sql"

// Open the bdd file
func OpenBDD() (*sql.DB, error) {
	return sql.Open("sqlite3", "structure/database.db")
}
