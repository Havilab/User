package model

type User struct {
	UserId       int32  `json:"userid"gorm:"primarykey"`
	UserFullName string `json:"userfullname"`
	Email        string `json:"email"`
	MobileNumber string `json:"mobilenumber"`
	Password     string `json:"password"`
	HouseNO      string `json:"houseno"`
	Area         string `json:"area"`
	City         string `json:"city"`
	State        string `json:"state"`
	PostalCode   string `json:"postalcode"`
	Country      string `json:"country"`
	CreateDate   string `json:"createdate"`
}

type Orders struct {
	OderId      int     `json:"orderid"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Image       string  `json:"image"`
	Userid      int32   `json:"userid"`
	TotalPrice  float32 `json:"totalprice"`
	CreateDate  string  `json:"createdate"`
}
