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

type All struct {
	Id        int    `gorm: "column:id"`
	FirstName string `gorm: "column:first_name"`
	LastName  string `gorm:"column:last_name"`
	Address   string `gorm:"column:address"`
	PhoneNumb string `gorm:"column:phone_numb"`
	Email     string `gorm:column:email`
	Pic       string `gorm:column:pic`
}
