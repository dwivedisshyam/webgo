package main

import (
	"net/http"

	"github.com/dwivedisshyam/webgo/pkg/webgo"
	"github.com/dwivedisshyam/webgo/pkg/webgo/types"
)

func main() {
	w := webgo.New()

	w.Get("/student", GetAll)
	w.Post("/student", Create)
	w.Start()
}

type Student struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func Create(c *webgo.Context) (interface{}, error) {
	s := Student{}

	c.Bind(&s)

	_, err := c.DB().Exec("insert into student values (?,?,?)", s.ID, s.Name, s.Age)
	if err != nil {
		return nil, &types.Error{Reason: err.Error()}
	}

	return &types.Response{Data: "Data", StatusCode: http.StatusCreated}, nil
}

func GetAll(c *webgo.Context) (interface{}, error) {
	res, _ := c.DB().Query("select * from student")

	students := make([]Student, 0)

	for res.Next() {
		s := Student{}
		res.Scan(&s.ID, &s.Name, &s.Age)
		students = append(students, s)
	}

	return students, nil
}
