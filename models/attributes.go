package models

import "gorm.io/gorm"

type AttributesJson struct {
	ID                 string `json:"id"`
	EventID            string `json:"event_id"`
	ObjectID           string `json:"object_id"`
	ObjectRelation     string `json:"object_relation"`
	Category           string `json:"category"`
	Type               string `json:"type"`
	Value1             string `json:"value1"`
	Value2             string `json:"value2"`
	ToIDS              bool   `json:"to_ids"`
	UUID               string `json:"uuid"`
	Timestamp          string `json:"timestamp"`
	Distribution       string `json:"distribution"`
	SharingGroupID     string `json:"sharing_group_id"`
	Comment            string `json:"comment"`
	Deleted            bool   `json:"deleted"`
	DisableCorrelation bool   `json:"disable_correlation"`
	FirstSeen          string `json:"first_seen"`
	LastSeen           string `json:"last_seen"`
	Value              string `json:"value"`
}

type AttributeDB struct {
	gorm.Model
	ItemID            string
	Category          string
	Type              string
	Value             string
	Deleted           bool
	SendToCorrelation bool
}

type AttributesCorrelation struct {
	ID       string `json:"id"`
	Category string `json:"category"`
	Type     string `json:"type"`
	Value    string `json:"value"`
}
