package router

import (
    "net/http"

    "github.com/jmoiron/sqlx"
    "pawntoqueen-go/internal/handler"
)

func SetupRouter(db *sqlx.DB) http.Handler {
    h := handler.NewHandler(db)

    mux := http.NewServeMux()
    mux.HandleFunc("/query", h.QueryHandler)
    return mux
}
