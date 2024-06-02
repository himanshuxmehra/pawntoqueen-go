package handler

import (
    "database/sql"
    "encoding/json"
    "net/http"

    "github.com/jmoiron/sqlx"
)

type Handler struct {
    db *sqlx.DB
}

func NewHandler(db *sqlx.DB) *Handler {
    return &Handler{db: db}
}

func (h *Handler) QueryHandler(w http.ResponseWriter, r *http.Request) {
    rows, err := h.db.Queryx("SELECT * FROM puzzles")
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var result []map[string]interface{}
    for rows.Next() {
        row := make(map[string]interface{})
        err = rows.MapScan(row)
        if err != nil {
            http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            return
        }
        result = append(result, row)
    }

    jsonResponse, err := json.Marshal(result)
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonResponse)
}
