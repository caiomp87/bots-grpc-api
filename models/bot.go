package models

import "time"

type Bot struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string    `json:"name,omitempty" bson:"name,omitempty"`
	Strategy  string    `json:"strategy,omitempty" bson:"strategy,omitempty"`
	UserID    string    `json:"userID,omitempty" bson:"userID,omitempty"`
	Squad     string    `json:"squad,omitempty" bson:"squad,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}
