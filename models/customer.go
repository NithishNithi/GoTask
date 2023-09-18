package models

type Customer struct {
	CustomerId  string     `json:"customerid" bson:"customerid"`
	Email       string     `json:"email" bson:"email"`
	FullName    string     `json:"fullname" bson:"fullname"`
	Password    string     `json:"password" bson:"password"`
	DateofBirth string     `json:"dateofbirth" bson:"dateofbirth"`
	PhoneNumber string     `json:"phonenumber" bson:"phonenumber"`
	Address     []*Address `json:"address" bson:"address"`
}

type Address struct {
	Country string `json:"country" bson:"country"`
	Street  string `json:"street" bson:"street"`
	City    string `json:"city" bson:"city"`
	State   string `json:"state" bson:"state"`
	Zip     string `json:"zip" bson:"zip"`
}

type CustomerResponse struct {
	CustomerId string `json:"customerid" bson:"customerid"`
}

// -----------

type Login struct {
	CustomerId string `json:"customerid" bson:"customerid"`
	Email      string `json:"email" bson:"email"`
	Password   string `json:"password" bson:"password"`
}

type Token struct {
	CustomerId string `json:"customerid" bson:"customerid"`
	Email      string `json:"email" bson:"email"`
	Token      string 	`json:"token" bson:"token"`
}
type TokenResponse struct {
	Token string `json:"token" bson:"token"`
}


