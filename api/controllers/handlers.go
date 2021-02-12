package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/astraker55/trade-marketing/api/models"
	"github.com/astraker55/trade-marketing/api/queries"
	"github.com/astraker55/trade-marketing/api/utils"
	"github.com/jmoiron/sqlx"
)

// Home wraps main page of API
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", "Welcome to StatInfoAPI")

}

// SaveStatHandler handles POST-requests
func (server *Server) SaveStatHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "POST" {
		http.Error(w, r.Method, http.StatusMethodNotAllowed)
		return
	}
	var event models.Event
	err := utils.DecodeJSONBody(w, r, &event)
	if err != nil {
		var mr *utils.MalformedRequest
		if errors.As(err, &mr) {
			http.Error(w, mr.Msg, mr.Status)
		} else {
			log.Println(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}
	_, err = event.SaveEvent(server.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	fmt.Fprintf(w, "Event was added: %+v", event)
}

// DropStatHandler handles request to clear db Table. Use it carefully
func (server *Server) DropStatHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, err := server.DB.Exec(queries.TruncateQuery)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	msg := []byte("Data was erased")
	w.Write(msg)
}

// GetStatHandler return sorted info about events
func (server *Server) GetStatHandler(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		from := r.URL.Query().Get("from")
		if len(from) < 1 {
			mr := utils.MalformedRequest{Status: http.StatusBadRequest, Msg: "Parameter from is empty"}
			http.Error(w, mr.Msg, mr.Status)
			return
		}
		f, err := time.Parse(models.Layout, from)
		if err != nil {
			mr := utils.MalformedRequest{Status: http.StatusBadRequest, Msg: "Wrong parameter from"}
			http.Error(w, mr.Msg, mr.Status)
			return
		}
		to := r.URL.Query().Get("to")
		if len(from) < 1 {
			mr := utils.MalformedRequest{Status: http.StatusBadRequest, Msg: "Parameter to is empty"}
			http.Error(w, mr.Msg, mr.Status)
			return
		}
		t, err := time.Parse(models.Layout, to)
		if err != nil {
			mr := utils.MalformedRequest{Status: http.StatusBadRequest, Msg: "Wrong parameter to"}
			http.Error(w, mr.Msg, mr.Status)
		}
		if t.Unix() <= f.Unix() {
			mr := utils.MalformedRequest{Status: http.StatusBadRequest, Msg: "Can't sort by these parameters"}
			http.Error(w, mr.Msg, mr.Status)
			return
		}
		var SortField string
		switch SortField := r.URL.Query().Get("sort"); SortField {
		case "Values":
			SortField = "values"
		case "Cost":
			SortField = "cost"
		case "Clicks":
			SortField = "clicks"
		default:
			SortField = "date"
		}
		var events models.EventsInfo
		err = db.Select(&events, queries.SelectQuery, f.Unix(), t.Unix(), SortField)
		if err != nil {
			http.Error(w, err.Error(), 404)
			return
		}
		if len(events) == 0 {
			w.Write([]byte("Empty repsonse"))
			return
		}
		eventsInfo, err := json.Marshal(events)
		if err != nil {
			http.Error(w, err.Error(), 405)
		}
		w.Write(eventsInfo)
	}
}
