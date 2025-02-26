<<<<<<< HEAD
<<<<<<< HEAD
package main

import (
  "encoding/json"
  "net/http"
  "time"
)
=======
type TimeResponse struct {
  Time string json:"time"
}
>>>>>>> 1a2c3a4 (Added feature Yarik)
=======
func timeHandler(w http.ResponseWriter, r *http.Request) {
response := TimeResponse{Time: time.Now().Format(time.RFC3339)}
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(response)
}
>>>>>>> bd328e9 (Updated by Lev)
