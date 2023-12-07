package main

import (
	"fmt"
	"net/http"
	"sync"
)

// DiskCount represents the structure of the dummy database entry for disk count.
type DiskCount struct {
	Count int
}

// AppDB is a dummy local database to store disk count.
var AppDB struct {
	sync.Mutex
	DiskCount DiskCount
}

func main() {
	// Initialize the dummy database with an initial disk count.
	AppDB.DiskCount = DiskCount{Count: 15}

	// Set up the API endpoint
	http.HandleFunc("/get-disk-count", getDiskCountHandler)

	// Start the server
	port := "8090"
	fmt.Printf("Server listening on :%s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Printf("Error starting the server: %s\n", err)
	}
}

func getDiskCountHandler(w http.ResponseWriter, r *http.Request) {
	// Assuming this endpoint only supports GET requests
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Retrieve the disk count from the dummy database
	AppDB.Lock()
	count := AppDB.DiskCount.Count
	AppDB.Unlock()

	// Send the disk count as a JSON response
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"disk_count": %d}`, count)
}
