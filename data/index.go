package data

import (
	"encoding/json"
	"io/ioutil"
)

type Role struct {
	Name   string
	Duties string
}

type User struct {
	Role Role
	Name string
}

type Company struct {
	Name  string
	Goals string
}

type Open struct {
	Role    Role
	Company Company
}

type Config struct {
	User User
	Open Open
}

func NewMock() Config {
	userRole := Role{"content specialist", "writing SQL queries for data engineers"}
	openRole := Role{"sql developer", "writing SQL queries and maintaining a database"}

	openCompany := Company{"Mock, Inc.", "delivering intuitive software solutions for non-trivial problems"}

	user := User{userRole, "Dev Developerson"}
	open := Open{openRole, openCompany}

	return Config{user, open}
}

func FromJson(path string) (Config, error) {
	var cfg Config
	raw, err := ioutil.ReadFile(path)

	if err != nil {
		return cfg, err
	}

	json.Unmarshal(raw, &cfg)

	return cfg, nil
}
