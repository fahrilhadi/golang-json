package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T)  {
	customer := Customer{
		FirstName: "Fahril",
		LastName: "Hadi",
		Age: 25,
		Married: false,
		Hobbies: []string{"Gaming", "Reading", "Coding"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T)  {
	jsonString := `{"FirstName":"Fahril","LastName":"Hadi","Age":25,"Married":false,"Hobbies":["Gaming","Reading","Coding"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Hobbies)
}

func TestJSONArrayComplex(t *testing.T)  {
	customer := Customer{
		FirstName: "Fahril",
		Addresses: []Address{
			{
				Street: "Jalan Belum Ada",
				Country: "Indonesia",
				PostalCode: "23234",
			},
			{
				Street: "Jalan Lagi Dibangun",
				Country: "Indonesia",
				PostalCode: "99999",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T)  {
	jsonString := `{"FirstName":"Fahril","LastName":"","Age":0,"Married":false,"Hobbies":null,"Addresses":[{"Street":"Jalan Belum Ada","Country":"Indonesia","PostalCode":"23234"},{"Street":"Jalan Lagi Dibangun","Country":"Indonesia","PostalCode":"99999"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Addresses)
}
 
func TestOnlyJSONArrayComplexDecode(t *testing.T)  {
	jsonString := `[{"Street":"Jalan Belum Ada","Country":"Indonesia","PostalCode":"23234"},{"Street":"Jalan Lagi Dibangun","Country":"Indonesia","PostalCode":"99999"}]`
	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}
	err := json.Unmarshal(jsonBytes, addresses)
	if err != nil {
		panic(err)
	}

	fmt.Println(addresses)
}

func TestOnlyJSONArrayComplex(t *testing.T)  {
	addresses := []Address{
		{
			Street: "Jalan Belum Ada",
			Country: "Indonesia",
			PostalCode: "23234",
		},
		{
			Street: "Jalan Lagi Dibangun",
			Country: "Indonesia",
			PostalCode: "99999",
		},
	}

	bytes, _ := json.Marshal(addresses)
	fmt.Println(string(bytes))
}