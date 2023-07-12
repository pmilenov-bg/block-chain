package models

type UserData struct {
	Data string
}

func (userData *UserData) ExportToString() string {
	return userData.Data
}
