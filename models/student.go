package models

import (
	"errors"
	"strconv"
	"time"
	"sync"
)

type Student struct {
	Id   string
	Name string
	Age  int
	Sex  string
}

var (
	StudentList sync.Map
)

func init() {
	s := Student{"1", "ben", 26, "male"}
	StudentList.Store("user_1", &s)
}

func AddStudent(s Student) string {
	s.Id = "user_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	StudentList.Store(s.Id, &s)
	return s.Id
}

func GetStudent(uid string) (s *Student, err error) {
	if s, ok := StudentList.Load(uid); ok {
		return s.(*Student), nil
	}
	return nil, errors.New("Student not exists")
}

func GetAllStudent() sync.Map {
	return StudentList
}

func UpdateStudent(uid string, ss *Student) (a *Student, err error) {
	if s, ok := StudentList.Load(uid); ok {
		if ss.Name != "" {
			s.(*Student).Name = ss.Name
		}
		if ss.Age != 0 {
			s.(*Student).Age = ss.Age
		}
		if ss.Sex != "" {
			s.(*Student).Sex = ss.Sex
		}

		return s.(*Student), nil
	}
	return nil, errors.New("User Not Exist")
}

func DeleteStudent(uid string) {
	StudentList.Delete(uid)
}
