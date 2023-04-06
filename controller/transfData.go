package controller

import "github.com/KbaYero/UTMStackDatasourcesMISP/models"

// AttrFromJsonToDB converts a json received from the api to an attribute to insert into the database
func AttrFromJsonToDB(attributeJson models.AttributesJson) models.AttributeDB {
	newDBAttr := models.AttributeDB{
		ItemID:   attributeJson.ID,
		Category: attributeJson.Category,
		Type:     attributeJson.Type,
		Value:    attributeJson.Value,
		Deleted:  attributeJson.Deleted,
	}
	return newDBAttr
}

// AttrFromDBCorrelation converts a database attribute into a json to send to the correlation
func AttrFromDBCorrelation(attrDB models.AttributeDB) models.AttributesCorrelation {
	newCorrelarionAttr := models.AttributesCorrelation{
		ID:       attrDB.ItemID,
		Category: attrDB.Category,
		Type:     attrDB.Type,
		Value:    attrDB.Value,
	}
	return newCorrelarionAttr
}
