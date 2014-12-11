package word

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB struct {
	*sqlx.DB
}

func (db *DB) Open() error {
	d, err := sqlx.Open("postgres", "timezone=UTC sslmode=disable")
	if err != nil {
		return err
	}

	db.DB = d

	return nil
}

func (db *DB) GetRandomWord() (*Word, error) {
	word := &Word{}
	err := db.Get(word, "select * from words offset random()*(select max(word_id) from words) limit 1")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("insert into word_history (word_id) values ($1)", word.Id)
	if err != nil {
		return nil, err
	}

	return word, nil
}

func (db *DB) GetRandom10() ([]*Word, error) {
	words := []*Word{}
	err := db.Select(&words, "select * from words order by random() limit 10")
	if err != nil {
		return nil, err
	}

	return words, nil
}

func (db *DB) GetHistory() ([]*Word, error) {
	words := []*Word{}
	err := db.Select(&words, "select w.* from words w inner join word_history wh on w.word_id = wh.word_id order by wh.created_at desc limit 10")
	if err != nil {
		return nil, err
	}

	return words, nil
}

func (db *DB) GetTop10() ([]*Word, error) {
	words := []*Word{}
	err := db.Select(&words, "select * from words order by rating desc limit 10")
	if err != nil {
		return nil, err
	}

	return words, nil
}

func (db *DB) GetBottom10() ([]*Word, error) {
	words := []*Word{}
	err := db.Select(&words, "select * from words order by rating asc limit 10")
	if err != nil {
		return nil, err
	}

	return words, nil
}

func (db *DB) UpVoteWord(word string, ip string) error {
	err := db.updateRating(word, 1, ip)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) DownVoteWord(word string, ip string) error {
	err := db.updateRating(word, -1, ip)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) updateRating(word string, vote int64, ip string) error {
	_, err := db.Exec("insert into votes (word, vote, ipaddress) values ($1, $2, $3)", word, vote, ip)
	if err != nil {
		return err
	}
	return nil
}
