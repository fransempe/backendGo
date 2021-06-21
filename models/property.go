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
	Google_Map_Data  GoogleMapData   `json:"google_map_data,omitempty"`
	City             City            `json:"city"`
	State            State           `json:"state"`
	Country          Country         `json:"country"`
	Neighborhood     string          `json:"neighborhood"`
	Rooms            int             `json:"rooms"`
	Bathrooms        int             `json:"bathrooms,omitempty"`
	Garages          int             `json:"garages,omitempty"`
	M2               int             `json:"m2"`
	M2_covered       int             `json:"m2_covered,omitempty"`
	Year             int             `json:"year,omitempty"`
	Price            int             `json:"price"`
	Amenities        []Amenities     `json:"amenities,omitempty"`
	Images           []Images        `json:"images,omitempty"`
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

type GoogleMapData struct {
	Lat_Lng string `json:"lat_lng"`
	Type    string `json:"type"`
	Zoom    int    `json:"zoom"`
	Exact   int    `json:"exact"`
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

type SearchParams struct {
	PropertyTypeSearch int    `json:"property_type"`
	TransactionType    int    `json:"transaction_type"`
	Text               string `json:"text"`
	Range_Min          int    `json:"range_min"`
	Range_Max          int    `json:"range_max"`
}

type Properties struct {
	Properties []Property `json:"properties"`
}
