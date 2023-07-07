package models

type DbKey struct {
	DbConnection struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		Dbname   string `json:"dbname"`
	}
}

type Config struct {
	LocalHost struct {
		Host string `json:"host"`
		Port string `json:"port"`
	}
}

type Name struct {
	Id        int    `gorm: "column:id"`
	FirstName string `gorm: "column:first_name"`
	LastName  string `gorm:"column:last_name"`
}