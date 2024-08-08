package db

//create table problem(
//    id int primary key auto_increment,
//    atcoderid varchar(20) ,
//    level int
//);

type Problem struct {
	Id        int    `json:"id"`
	Atcoderid string `json:"atcoderid"`
	Level     int    `json:"level"`
}

func (p Problem) TableName() string {
	return "problem"
}

//create view problem_view as
//select user.id as userid , problem.id as problemid ,problem.atcoderid
//from user,problem
//where user.level = problem.level ;

type ProblemView struct {
	Userid    int    `json:"userid"`
	Truename  string `json:"truename"`
	Username  string `json:"username"`
	Problemid int    `json:"problemid"`
	Atcoderid string `json:"atcoderid"`
	Level     int    `json:"level"`
}

func (pv ProblemView) TableName() string {
	return "problem_view"
}

// create view problem_ac_view as
// SELECT problem_view.*,
//
//	COUNT(submit.id) AS account
//
// FROM problem_view
//
//	     LEFT JOIN submit ON (
//	submit.statuscanonical = 'AC' AND
//	submit.atcoderid = problem_view.atcoderid AND
//	submit.userid = problem_view.userid
//	)
//
// GROUP BY problem_view.problemid,problem_view.userid;
type ProblemACView struct {
	Userid    int    `json:"userid"`
	Truename  string `json:"truename"`
	Username  string `json:"username"`
	Problemid int    `json:"problemid"`
	Atcoderid string `json:"atcoderid"`
	Account   int    `json:"account"`
	Level     int    `json:"level"`
}

func (pav ProblemACView) TableName() string {
	return "problem_ac_view"
}

type ProblemCountView struct {
	Userid    int    `json:"userid"`
	Truename  string `json:"truename"`
	Username  string `json:"username"`
	Problemid int    `json:"problemid"`
	Atcoderid string `json:"atcoderid"`
	Account   int    `json:"account"`
	Wacount   int    `json:"wacount"`
	Level     int    `json:"level"`
}

func (pcv ProblemCountView) TableName() string {
	return "problem_count_view"
}
