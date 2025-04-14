package golang_json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestEncoder(t *testing.T)  {
	writer, _ := os.Create("CustomerOut.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName: "Fahril",
		LastName: "Hadi",
	}

	encoder.Encode(customer)
	fmt.Println(customer)
}