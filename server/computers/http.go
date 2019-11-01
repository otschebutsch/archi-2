package computers

import (
	//"encoding/json"
	"test/server/tools"
	"log"
	"net/http"
)

// Computers HTTP handler.
type HttpHandlerFunc http.HandlerFunc

// HttpHandler creates a new instance of computers HTTP handler.
func HttpHandler(store *Store) HttpHandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handleListComputers(store, rw)
			handleSetDisk(rw, store)
			//handleListDisks(store, rw)
		} else if r.Method == "POST" {
			//handleComputerCreate(r, rw, store)
			//handleDiskCreate(r, rw, store)
			//handleSetDisk(r, rw, store)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}


func handleListComputers(store *Store, rw http.ResponseWriter) {
	res, err := store.ListComputers()
	if err != nil {
		log.Printf("Error making query to the db: %s", err)
		tools.WriteJsonInternalError(rw)
		return
	}
	tools.WriteJsonOk(rw, res)
}

func handleSetDisk(rw http.ResponseWriter, store *Store) {
	res, err := store.SetDisk()
	if err != nil {
		log.Printf("Error making query to the db: %s", err)
		tools.WriteJsonInternalError(rw)
		return
	}
	tools.WriteJsonOk(rw, res)
}
/*
func handleComputerCreate(r *http.Request, rw http.ResponseWriter, store *Store) {
	var c Computer
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		log.Printf("Error decoding computer input: %s", err)
		tools.WriteJsonBadRequest(rw, "bad JSON payload")
		return
	}
	err := store.CreateComputer(c.Name, c.CpuCount)
	if err == nil {
		tools.WriteJsonOk(rw, &c)
	} else {
		log.Printf("Error inserting record: %s", err)
		tools.WriteJsonInternalError(rw)
	}
}

func handleDiskCreate(r *http.Request, rw http.ResponseWriter, store *Store) {
	var c Disk
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		log.Printf("Error decoding disk input: %s", err)
		tools.WriteJsonBadRequest(rw, "bad JSON payload")
		return
	}
	err := store.CreateDisk(c.DiskSpace, c.Id_c)
	if err == nil {
		tools.WriteJsonOk(rw, &c)
	} else {
		log.Printf("Error inserting record: %s", err)
		tools.WriteJsonInternalError(rw)
	}
}

func handleListDisks(store *Store, rw http.ResponseWriter) {
	res, err := store.ListDisks()
	if err != nil {
		log.Printf("Error making query to the db: %s", err)
		tools.WriteJsonInternalError(rw)
		return
	}
	tools.WriteJsonOk(rw, res)
}
*/
