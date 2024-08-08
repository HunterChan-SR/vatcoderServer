package db

// create table submit(
// id int primary key,
// statuscanonical varchar(20),
// userid int,
// atcoderid varchar(20)
// );
type Submit struct {
	Id              int    `json:"id"`
	Statuscanonical string `json:"statuscanonical"`
	Userid          int    `json:"userid"`
	Atcoderid       string `json:"atcoderid"`
}

func (s Submit) TableName() string {
	return "submit"
}
