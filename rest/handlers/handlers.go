package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"sodaCounter_rest/pkg/db"
	"sodaCounter_rest/pkg/framework/request"
	"sodaCounter_rest/pkg/framework/response"

	"github.com/k0kubun/pp"
)

// Handler struct to hold handler methods
type Handler struct {
	DB *sql.DB
}

// CreateMaterialInput represents the expected input for creating a material
type CreateMaterialInput struct {
	ID     int    `json:"id"`
	Soda   string `json:"soda"`
	Amount uint   `json:"amount"`
}

// NewHandler creates a new Handler instance
func NewHandler(db *sql.DB) *Handler {
	return &Handler{DB: db}
}

// GetAllMaterials handles the retrieval of all materials
func (h *Handler) GetAllSoda(w http.ResponseWriter, r *http.Request) {
	rows, err := h.DB.Query("SELECT id, Soda, Amount, author FROM materials")
	if err != nil {
		response.Err(w, response.ErrInternal)
		return
	}
	defer rows.Close()
}

// CreateMaterial handles the creation of a new material
func (h *Handler) CreateSoda(w http.ResponseWriter, r *http.Request, tableName string) {
	data := CreateMaterialInput{} // Initialize an empty struct to hold the input data

	// Decode the JSON request body into the data-struct
	if err := request.JSON(r.Body, &data); err != nil {
		response.Err(w, response.ErrInvalidInput)
		log.Printf("error decoding request body: %v", err)
		return
	}

	// For debugging purposes, pretty-print the received data to the console
	pp.Print(data)

	// TODO: Add logic to save the new material to the database (or similar datasource)
	err := db.TableAddSoda(h.DB, tableName, data.Soda, data.Amount)
	if err != nil {
		response.Err(w, response.ErrInternal)
		log.Printf("error saving material to database: %v", err)
		return
	}

	// Respond with a success message
	w.Write([]byte("Successfully registered material")) // This should write back success message or created material ID etc.
}

// GetMaterial handles the retrieval of a specific material
func (h *Handler) GetSoda(w http.ResponseWriter, r *http.Request, tableName string, id int) {

	// Query the database for the material with the specified ID
	rows, err := db.TableGetSoda(h.DB, tableName, id)
	if err != nil {
		response.Err(w, response.ErrInternal)
		log.Printf("error retrieving material from database: %v", err)
		return
	}

	var material CreateMaterialInput // Struct to hold the retrieved material data, based in CreateMaterialInput for standardization

	// Iterate over the result set and scan the data into the material struct
	if rows.Next() {
		if err := rows.Scan(&material.ID, &material.Soda, &material.Amount); err != nil {
			response.Err(w, response.ErrInternal)
			return
		}
	} else {
		response.Err(w, response.ErrNotFound)
	}

	// Ensure rows are closed after processing
	defer rows.Close()

	response.JSON(w, http.StatusOK, material) // Respond with the retrieved material as JSON

	pp.Printf("%+v", material)
}

// UpdateMaterial handles the updating of a specific material
func (h *Handler) UpdateSoda(w http.ResponseWriter, r *http.Request, tableName string, id int) {
	// Attempt to decode the request body into a CreateMaterialInput struct
	data := CreateMaterialInput{}
	if err := request.JSON(r.Body, &data); err != nil {
		response.Err(w, response.ErrInvalidInput)
		log.Printf("error decoding request body: %v", err)
		return
	}

	// Update the material in the database
	err := db.TableUpdateSoda(h.DB, tableName, id, data.Soda, data.Amount)
	if err != nil {
		response.Err(w, response.ErrInternal)
		log.Printf("error updating material in database: %v", err)
		return
	}

	// Respond with a success message
	response.Success(w, response.SuccessUpdated)

	// For debugging purposes, pretty-print the updated data to the console
	//pp.Printf("Updated material with ID: %d", id)
}

// DeleteMaterial handles the deletion of a specific material
func (h *Handler) DeleteSoda(w http.ResponseWriter, r *http.Request, tableName string, id int) {
	w.Write([]byte("Delete material"))
}
