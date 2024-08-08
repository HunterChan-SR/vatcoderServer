package db

//create table level(
//    id int primary key auto_increment,
//    title varchar(20)
//);

type Level struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

func (l Level) TableName() string {
	return "level"
}
