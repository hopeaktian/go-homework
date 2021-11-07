package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

type User struct {
	id int
	name string
}

func QueryUser(userId int) (User, error) {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/golearn")
	if err != nil {
		return User{}, errors.Wrap(err, "")
	}

	var (
		name string
		id	int
	)
	err = db.QueryRow("select id, name from users where id = ?", userId).Scan(&id, &name)
	if errors.Cause(err) == sql.ErrNoRows {
		return User{}, errors.Wrap(err, fmt.Sprintf("id = %d", userId))
	}

	if err != nil {
		// unknown errors
		return User{}, errors.Wrap(err, "")
	}
	return User{id: id, name: name}, nil

}

func main() {
	user, err := QueryUser(2)
	if errors.Cause(err) == sql.ErrNoRows {
		fmt.Println("query no data, detail errors: ", err)
		return
	}
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}
	fmt.Printf("query ok: %#v\n", user)

	for {
		fmt.Println("a")
	}
}
