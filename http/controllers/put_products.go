package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/fusionxx23/ecommerce-go/http/models"
)

type ThumbnailUpdateRequest struct {
	ThumbnailIDOne string `json:"thumbnailIdOne,omitempty"`
	ThumbnailIDTwo string `json:"thumbnailIdTwo,omitempty"`
	ProductID      string `json:"productId"`
}

func updateThumbnails(w http.ResponseWriter, r *http.Request) {
	var payload ThumbnailUpdateRequest
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
	}
	p, err := strconv.ParseInt(payload.ThumbnailIDOne, 10, 64)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
	}
	if payload.ThumbnailIDOne != "" {
		t, err := strconv.ParseInt(payload.ThumbnailIDOne, 10, 64)
		if err != nil {
			http.Error(w, "Invalid Thumbnail ID", http.StatusBadRequest)
		}
		_, err = models.GetProductImage(t)
		if err != nil {
			http.Error(w, "Thumbnail not found", http.StatusNotFound)
		} else {
			models.UpdateThumbnailOne(p, t)
		}
	}

	if payload.ThumbnailIDTwo != "" {
		t, err := strconv.ParseInt(payload.ThumbnailIDTwo, 10, 64)
		if err != nil {
			http.Error(w, "Invalid Thumbnail ID", http.StatusBadRequest)
		}
		_, err = models.GetProductImage(t)
		if err != nil {
			http.Error(w, "Thumbnail not found", http.StatusNotFound)
		} else {
			models.UpdateThumbnailTwo(p, t)
		}
	}
}
