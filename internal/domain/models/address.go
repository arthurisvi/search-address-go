package models

import "fmt"

type AddressModel struct {
	Address1      string   `json:"address1" xml:"address1"`
	Address2      string   `json:"address2" xml:"address2"`
	Neighborhood  string   `json:"neighborhood" xml:"neighborhood"`
	AddressNumber int      `json:"addressNumber" xml:"addressNumber"`
	ZipCode       string   `json:"zipCode" xml:"zipCode"`
	IbgeCode      string   `json:"ibgeCode" xml:"ibgeCode"`
	Locality      Locality `json:"locality" xml:"locality"`
}

type Locality struct {
	City    string `json:"city" xml:"city"`
	State   string `json:"state" xml:"state"`
	Country string `json:"country" xml:"country"`
}

func (a AddressModel) String() string {
	if a.Address2 == "" {
		return fmt.Sprintf("%v, %v, %v-%v, CEP: %v", a.Address1, a.AddressNumber, a.Locality.City, a.Locality.State, a.ZipCode)
	}
	return fmt.Sprintf("%v, %v, %v, %v-%v, CEP: %v", a.Address1, a.AddressNumber, a.Address2, a.Locality.City, a.Locality.State, a.ZipCode)
}
