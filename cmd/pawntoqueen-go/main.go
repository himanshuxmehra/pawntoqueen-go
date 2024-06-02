package main

import (
    "log"
    "net/http"
    "os"
    "pawntoqueen-go/internal/config"
    "pawntoqueen-go/internal/router"
    "pawntoqueen-go/pkg/logger"
)

func main() {
    // Initialize configuration
    config.Load()

    // Initialize logger
    logger.Init()

    // Initialize database
    db, err := config.InitDB()
    if err != nil {
        log.Fatalf("Failed to initialize database: %v", err)
    }
    defer db.Close()

    // Setup router
    r := router.SetupRouter(db)

    // Start the server
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    log.Printf("Starting server on port %s", port)
    log.Fatal(http.ListenAndServe(":"+port, r))
}
