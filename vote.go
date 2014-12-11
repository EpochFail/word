package word

import "time"

type Vote struct {
	Id        int64 `db:"vote_id"`
	Word      string
	Vote      int64
	CreatedAt time.Time `db:"created_at"`
	IPAddress string
}
