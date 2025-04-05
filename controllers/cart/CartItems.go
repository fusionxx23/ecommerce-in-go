package cart

import "net/http"

func CartItem(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		// get product quanity
		// check if user already has item in cart
		//  check if enough product quanity in stock
		// if so increase the quantity
	}
	if r.Method == "DELETE" {

	}
}
