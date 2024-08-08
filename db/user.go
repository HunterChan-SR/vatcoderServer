package db

//create table user(
//    id int primary key auto_increment,
//    username varchar(20) ,
//    password varchar(20) ,
//    truename varchar(20) ,
//    level int
//);

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Truename string `json:"truename"`
	Level    int    `json:"level"`
}

func (u User) TableName() string {
	return "user"
}

//create view user_view as
//select user.id as id,username,truename,level.title as level
//from user,level
//where user.level = level.id;

type UserView struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Truename string `json:"truename"`
	Level    string `json:"level"`
}

func (uv UserView) TableName() string {
	return "user_view"
}
