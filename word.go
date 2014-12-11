package word

type Word struct {
	Id     int64  `json:"id" db:"word_id"`
	Word   string `json:"word"`
	Rating int64  `json:"rating"`
}
