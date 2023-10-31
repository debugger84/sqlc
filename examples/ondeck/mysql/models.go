// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package ondeck

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"
)

type VenueStatus string

const (
	VenueStatusOpen   VenueStatus = "open"
	VenueStatusClosed VenueStatus = "closed"
)

func (e *VenueStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = VenueStatus(s)
	case string:
		*e = VenueStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for VenueStatus: %T", src)
	}
	return nil
}

type NullVenueStatus struct {
	VenueStatus VenueStatus `json:"venue_status"`
	Valid       bool        `json:"valid"` // Valid is true if VenueStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullVenueStatus) Scan(value interface{}) error {
	if value == nil {
		ns.VenueStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.VenueStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullVenueStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.VenueStatus), nil
}

type City struct {
	Slug string `json:"slug"`
	Name string `json:"name"`
}

// Venues are places where muisc happens
type Venue struct {
	ID uint64 `json:"id"`
	// Venues can be either open or closed
	Status   VenueStatus    `json:"status"`
	Statuses sql.NullString `json:"statuses"`
	// This value appears in public URLs
	Slug            string         `json:"slug"`
	Name            string         `json:"name"`
	City            string         `json:"city"`
	SpotifyPlaylist string         `json:"spotify_playlist"`
	SongkickID      sql.NullString `json:"songkick_id"`
	Tags            sql.NullString `json:"tags"`
	CreatedAt       time.Time      `json:"created_at"`
}
