// Package definitions godoc
// api requests and responses types
// with swagger godoc attributes
// and json un/marshaling attributes
// request types, to parse http/user inputs
// response types, to wrap data for http.ResponseWriter
package definitions

import "time"

// GormFields gorm.Model alias
// to have `time related` sql fields with proper json attributes
type GormFields struct {
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
}
