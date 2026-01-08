package models

type Name struct {
	First string
	Last  string
}

type Group string

const (
	IT      Group = "IT"
	HR      Group = "HR"
	PR      Group = "PR"
	IO      Group = "IO"
	JFR     Group = "JFR"
	GRAFIKA Group = "Grafika"
)

type User struct {
	ID    uint8
	Name  Name
	Email string
	Group Group
}
