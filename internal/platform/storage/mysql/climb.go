package mysql

const CLIMB_TABLE = "climbs"

type sqlClimb struct {
	Id          string `db:"id"`
	Date        string `db:"date"`
	Grade       string `db:"grade"`
	Description string `db:"description"`
	Area        string `db:"area"`
}
