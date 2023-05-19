package main

import (
	"os"
	"text/template"
)


type (
	Location struct {
		Street  string
		ZipCode string
	}

	User struct {
		Username  string
		Locations map[string]Location
	}

	UsersPage struct {
		Title string
		Users []User
	}
)

func main() {

	message, err := os.ReadFile("./example.gohtml")
	if err != nil {
		panic(err)
	}

	t, err := template.New("UsersPage").Parse(string(message))
	if err != nil {
		panic(err)
	}

	p := UsersPage{
		Title: "测试语法",
		Users: []User{
			{
				Username: "johnsmith",
				Locations: map[string]Location{
					"Home": {
						Street:  "Main Street",
						ZipCode: "20192",
					},
				},
			},
		},
	}

	err = t.Execute(os.Stdout, p)
	if err != nil {
		panic(err)
	}
}
