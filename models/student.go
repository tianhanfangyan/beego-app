package models

import (
	"errors"
	"strconv"
	"time"
)

type Student struct {
	Id   string
	Name string
	Age  int
	Sex  string
}

var (
	StudentList map[string]*Student
)

func init() {
	StudentList = make(map[string]*Student)
	s := Student{"1", "ben", 26, "male"}
	StudentList["user_1"] = &s
}

func AddStudent(s Student) string {
	s.Id = "user_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	StudentList[s.Id] = &s
	return s.Id
}

func GetStudent(uid string) (s *Student, err error) {
	if s, ok := StudentList[uid]; ok {
		return s, nil
	}
	return nil, errors.New("Student not exists")
}

func GetAllStudent() map[string]*Student {
	return StudentList
}

func UpdateStudent(uid string, ss *Student) (a *Student, err error) {
	if s, ok := StudentList[uid]; ok {
		if ss.Name != "" {
			s.Name = ss.Name
		}
		if ss.Age != 0 {
			s.Age = ss.Age
		}
		if ss.Sex != "" {
			s.Sex = ss.Sex
		}

		return s, nil
	}
	return nil, errors.New("User Not Exist")
}

func DeleteStudent(uid string) {
	delete(StudentList, uid)
}
