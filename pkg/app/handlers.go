package app

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/christianh814/golist-api/pkg/db"
	"github.com/christianh814/golist-api/pkg/entities"
	"github.com/gorilla/mux"
)

// CreateProduct creates a new product in the DB
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	// Set the header to JSON
	w.Header().Set("Content-Type", "application/json")

	// Defines a new product variable.
	var product entities.Product

	// Decodes the Body of the Incoming JSON request and maps it to the newly created product variable.
	json.NewDecoder(r.Body).Decode(&product)

	// Using GORM, we try to create a new product by passing in the parsed product. This would ideally create a new record in the products table for us.
	db.Instance.Create(&product)

	// Returns the newly created product data back to the client.
	json.NewEncoder(w).Encode(product)
}

// checkIfProductExists checks if a product exists in the DB
func checkIfProductExists(productId string) bool {
	// Defines a new product variable.
	var product entities.Product

	// We try to find a product by its ID. This would ideally find a record in the products table for us.
	db.Instance.First(&product, productId)

	// If the product is found, return true.
	if product.ID == 0 {
		return false
	}
	return true
}

// GetProductById gets a product by its ID
func GetProductById(w http.ResponseWriter, r *http.Request) {
	// Gets the Product Id from the Query string of the request.
	productId := mux.Vars(r)["id"]

	// If the id is not found in the product table, send "Product Not Found!"
	if checkIfProductExists(productId) == false {
		json.NewEncoder(w).Encode("Product Not Found!")
		return
	}

	// the product table is queried with the product Id. This would fill in the product details to the newly created product variable.
	var product entities.Product
	db.Instance.First(&product, productId)

	//  encode the product variable and send it back to the client.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

// GetProducts gets all products from the DB
func GetProducts(w http.ResponseWriter, r *http.Request) {
	// define an empty new list of products.
	var products []entities.Product

	// Maps all the available products into the product list variable.
	db.Instance.Find(&products)

	// encode the products variable and returns it back to the client.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

//  UpdateProduct updates a product in the DB
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	// MUX extracts the id from the URL and assigns the value to the id variable
	productId := mux.Vars(r)["id"]

	// checks if the passed product Id actually exists in the product table
	if checkIfProductExists(productId) == false {
		json.NewEncoder(w).Encode("Product Not Found!")
		return
	}

	// queries the product record to the product variable.
	var product entities.Product

	// JSON decoder then converts the request body to a product variable
	db.Instance.First(&product, productId)
	json.NewDecoder(r.Body).Decode(&product)
	db.Instance.Save(&product)

	// encode the product variable and send it back to the client.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

// DeleteProduct deletes a product from the DB
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	// Set the header type
	w.Header().Set("Content-Type", "application/json")

	// Extracts the id to be deleted from the request URL.
	productId := mux.Vars(r)["id"]
	// Checks if the ID is actually available in the product table.
	if checkIfProductExists(productId) == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Product Not Found!")
		return
	}

	// Then we create a new product variable.
	var product entities.Product

	// deletes the product by ID
	db.Instance.Delete(&product, productId)

	// encode the product variable and send it back to the client.
	json.NewEncoder(w).Encode("Product Deleted Successfully!")
}

// Return OK status
func ReturnHealth(w http.ResponseWriter, r *http.Request) {
	// Write the status code
	w.WriteHeader(200)

	// set header type
	w.Header().Set("Content-Type", "application/json")

	// write the message
	resp := make(map[string]string)
	resp["status"] = "OK"

	// encode the response report any errors
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	// write response
	w.Write(jsonResp)
}
