package computers

import (
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
