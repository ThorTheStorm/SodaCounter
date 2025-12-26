package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"sodaCounter_rest/pkg/db"
	"sodaCounter_rest/pkg/framework/request"
	"sodaCounter_rest/pkg/framework/response"
	"sodaCounter_rest/pkg/logging"
)

// Handler struct to hold handler methods
type Handler struct {
	DB        *sql.DB
	TableName string
}

// CreateSodaInput represents the expected input for creating a soda
type CreateSodaInput struct {
	ID           int    `json:"id"`
	Soda         string `json:"soda"`
	Amount       uint   `json:"amount"`
	Manufacturer string `json:"manufacturer,omitempty"` // Optional field, but strongly recommended
}

type CreateSodaBatchInput struct {
	Sodas []CreateSodaInput `json:"sodas"`
}

// NewHandler creates a new Handler instance
func NewHandler(db *sql.DB, tableName string) *Handler {
	return &Handler{
		DB:        db,
		TableName: tableName,
	}
}

// GetAllsodas handles the retrieval of all sodas
func (h *Handler) GetAllSoda(w http.ResponseWriter, r *http.Request) {
	rows, err := h.DB.Query(fmt.Sprintf("SELECT * FROM %s", h.TableName))
	if err != nil {
		response.Err(w, response.ErrInternal)
		return
	}
	defer rows.Close()

	//var sodaBatch CreateSodaBatchInput // Struct to hold the retrieved soda data, based in CreateSodaInput for standardization
	var sodas []CreateSodaInput

	// Iterate over the result set and scan the data into the soda struct
	for rows.Next() {
		var sodaItem CreateSodaInput
		if err := rows.Scan(&sodaItem.ID, &sodaItem.Soda, &sodaItem.Amount, &sodaItem.Manufacturer); err != nil {
			response.Err(w, response.ErrInternal)
			return
		}
		sodas = append(sodas, sodaItem)
		// } else {
		// 	response.Err(w, response.ErrNotFound)
	}

	if err := rows.Err(); err != nil {
		response.Err(w, response.ErrInternal)
		logging.AppLog(logging.ErrorLog, fmt.Sprintf("error iterating over rows: %v", err))
		return
	}

	response.JSON(w, http.StatusOK, sodas)
}

// Createsoda handles the creation of a new soda
func (h *Handler) CreateSoda(w http.ResponseWriter, r *http.Request, tableName string) {
	data := CreateSodaInput{} // Initialize an empty struct to hold the input data

	// Decode the JSON request body into the data-struct
	if err := request.JSON(r.Body, &data); err != nil {
		response.Err(w, response.ErrInvalidInput)
		logging.AppLog(logging.ErrorLog, fmt.Sprintf("error decoding request body: %v", err))
		return
	}

	// For debugging purposes, pretty-print the received data to the console
	// pp.Print(data)

	// TODO: Add logic to save the new soda to the database (or similar datasource)
	dbReturn, err := db.TableAddSoda(h.DB, tableName, data.Soda, data.Amount, data.Manufacturer)
	if err != nil {
		response.Err(w, response.ErrInternal)
		logging.AppLog(logging.ErrorLog, fmt.Sprintf("error saving soda to database: %v", err))
		return
	}

	data.ID = dbReturn // Assign the returned ID from the database to the data struct

	// Respond with a success message
	response.JSON(w, http.StatusCreated, data) // This should write back success message or created soda ID etc.
}

// Getsoda handles the retrieval of a specific soda
func (h *Handler) GetSoda(w http.ResponseWriter, r *http.Request, tableName string, id int) {

	// Query the database for the soda with the specified ID
	rows, err := db.TableGetSoda(h.DB, tableName, id)
	if err != nil {
		response.Err(w, response.ErrInternal)
		logging.AppLog(logging.ErrorLog, fmt.Sprintf("error retrieving soda from database: %v", err))
		return
	}

	var soda CreateSodaInput // Struct to hold the retrieved soda data, based in CreateSodaInput for standardization

	// Iterate over the result set and scan the data into the soda struct
	if rows.Next() {
		if err := rows.Scan(&soda.ID, &soda.Soda, &soda.Amount); err != nil {
			response.Err(w, response.ErrInternal)
			return
		}
	} else {
		response.Err(w, response.ErrNotFound)
		return
	}

	// Ensure rows are closed after processing
	defer rows.Close()

	response.JSON(w, http.StatusOK, soda) // Respond with the retrieved soda as JSON
}

// Updatesoda handles the updating of a specific soda
func (h *Handler) UpdateSoda(w http.ResponseWriter, r *http.Request, tableName string, id int) {
	// Attempt to decode the request body into a CreateSodaInput struct
	data := CreateSodaInput{}
	if err := request.JSON(r.Body, &data); err != nil {
		response.Err(w, response.ErrInvalidInput)
		logging.AppLog(logging.ErrorLog, fmt.Sprintf("error decoding request body: %v", err))
		return
	}

	// Update the soda in the database
	err := db.TableUpdateSodaAmount(h.DB, tableName, id, data.Amount)
	if err != nil {
		response.Err(w, response.ErrInternal)
		logging.AppLog(logging.ErrorLog, fmt.Sprintf("error updating soda in database: %v", err))
		return
	}

	// Respond with a success message
	response.Success(w, response.SuccessUpdated)

	// For debugging purposes, pretty-print the updated data to the console
	//pp.Printf("Updated soda with ID: %d", id)
}

// Deletesoda handles the deletion of a specific soda
func (h *Handler) DeleteSoda(w http.ResponseWriter, r *http.Request, tableName string, id int) {
	err := db.TableDeleteSoda(h.DB, tableName, id)
	if err != nil {
		response.Err(w, response.ErrInternal)
		logging.AppLog(logging.ErrorLog, fmt.Sprintf("error deleting soda from database: %v", err))
		return
	}

	// Respond with a success message
	response.Success(w, response.SuccessDeleted)
}
