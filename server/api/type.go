package api

type User struct {
    UserID string `json:"userId"`
}

type Shop struct {
	ID   string `bson:"_id"`
	shopName string `bson:"shopName"`
	phone string `bson:"phone"`
	address string `bson:"address"`
}

type Hairdresser struct {
	ID        string `bson:"_id"`
	ShopID    string `bson:"shop_id"`
	FirstName string `bson:"firstName"`
	LastName  string `bson:"lastName"`
}

type LoginRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

type RegistrationRequest struct {
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	IsShop    bool    `json:"is_shop"`
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
	ShopName  *string `json:"shop_name,omitempty"`
	Phone     *string `json:"phone,omitempty"`
	Address   *string `json:"address,omitempty"`
}

type Reservation struct {
	ID         string `json:"_id" bson:"_id"`
	ShopID     string `json:"shop_id" bson:"shop_id"`
	EmployeeID string `json:"employee_id" bson:"employee_id"`
	UserID     string `json:"user_id" bson:"user_id"`
	Date       string `json:"date" bson:"date"`
	Hours      string `json:"hours" bson:"hours"`
}
