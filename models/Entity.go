package models

// User struct which contains a name
// a type and a list of social links
type Entity struct {
	ID               int             `json:"id"`
	Title            string          `json:"title"`
	Property_Type    PropertyType    `json:"property_type"`
	Transaction_Type TransactionType `json:"transaction_type"`
	Currency         Currency        `json:"currency"`
	Address          string          `json:"address"`
	Address_Number   string          `json:"address_number"`
	Status           string          `json:"status"`
}

// PropertyType struct
type PropertyType struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// TransactionType struct
type TransactionType struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// Currency struct
type Currency struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var Entities []Entity
