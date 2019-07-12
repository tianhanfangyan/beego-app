package models

import (
	"errors"
	"strconv"
	"sync"
	"time"
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

func GetAllStudent() map[string]*Student {
	Datalist := make(map[string]*Student)

	//Range
	//遍历sync.Map, 要求输入一个func作为参数
	f := func(k, v interface{}) bool {
		//这个函数的入参、出参的类型都已经固定，不能修改
		//可以在函数体内编写自己的代码，调用map中的k,v
		Datalist[k.(string)] = v.(*Student)
		return true
	}
	StudentList.Range(f)

	return Datalist
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
