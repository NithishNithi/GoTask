package models

type Customer struct {
	CustomerId  string `json:"customerid" bson:"customerid"`
	Email       string `json:"email" bson:"email"`
	FullName    string `json:"fullname" bson:"fullname"`
	Password    string `json:"password" bson:"password"`
	DateofBirth string `json:"dateofbirth" bson:"dateofbirth"`
	PhoneNumber string `json:"phonenumber" bson:"phonenumber"`
	HouseNo     string `json:"houseno" bson:"houseno"`
	Street      string `json:"street" bson:"street"`
	City        string `json:"city" bson:"city"`
	Country     string `json:"country" bson:"country"`
	Zip         string `json:"zip" bson:"zip"`
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
	Token      string `json:"token" bson:"token"`
}
type TokenResponse struct {
	Token string `json:"token" bson:"token"`
}

type CreateCustomerResponse struct {
	Message    string `json:"message"`
}
