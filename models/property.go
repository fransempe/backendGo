package models

// User struct which contains a name
// a type and a list of social links
type Property struct {
	ID               int             `json:"id"`
	Title            string          `json:"title"`
	Property_Type    PropertyType    `json:"property_type"`
	Transaction_Type TransactionType `json:"transaction_type"`
	Currency         Currency        `json:"currency"`
	Address          string          `json:"address"`
	Address_Number   string          `json:"address_number"`
	City             City            `json:"city"`
	State            State           `json:"State"`
	Country          Country         `json:"Country"`
	Neighborhood     string          `json:"neighborhood"`
	Rooms            int             `json:"rooms"`
	M2               int             `json:"m2"`
	Status           string          `json:"status"`
	Payment          []string        `json:"payment"`
	Disposition      []string        `json:"disposition"`
	Tags             []string        `json:"tags"`
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

type City struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type State struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Country struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Amenities struct {
	Code  string `json:"code"`
	Name  string `json:"name"`
	Group string `json:"group"`
}

type Images struct {
	Original string `json:"original"`
	Small    string `json:"small"`
	Medium   string `json:"medium"`
	Big      string `json:"big"`
}

type Properties struct {
	Properties []Property `json:"properties"`
}
