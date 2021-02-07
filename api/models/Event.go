package models

import (
	"github.com/astraker55/trade-marketing/api/queries"
	"github.com/jmoiron/sqlx"
)

// Event is a main struct of app
type Event struct {
	Date   *CustomDate `json:"date" db:"date"`
	Views  int         `json:"views,omitempy" db:"views"`
	Clicks int         `json:"clicks,omitempy" db:"clicks"`
	Cost   float32     `json:"cost,omitempy" db:"cost"`
	Cpc    float32     `json:"cpc,omitempy" db:"cpm"`
	Cpm    float32     `json:"cpm,omitempy" db:"cpc"`
}

// SaveEvent put info into db
func (e *Event) SaveEvent(db *sqlx.DB) (*Event, error) {

	var err error
	_, err = db.NamedExec(queries.InsertQuery, &e)
	if err != nil {
		return &Event{}, err
	}
	return e, nil
}

// EventsInfo contains objects of Event
type EventsInfo []Event
