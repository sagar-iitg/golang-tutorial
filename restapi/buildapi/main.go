package main

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

//Model for Course file

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"courseprice"`
	Author      *Author `json:"author"`
}

//Slice of courses -- fake database
var courses []Course

//middleware,helper file
func (c *Course) IsEmpty() bool {

	return c.CourseId == "" && c.CourseName == ""

}
func main() {

}
