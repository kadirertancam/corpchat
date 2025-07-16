package chat

import "github.com/jmoiron/sqlx"

func SaveMessage(db *sqlx.DB, m Message) error {
	_, err := db.Exec(`INSERT INTO messages(from_uid,to_uid,body,created_at)
		VALUES($1,$2,$3,now())`, m.FromUID, m.ToUID, m.Body)
	return err
}