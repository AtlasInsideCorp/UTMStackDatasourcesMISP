package controller

import "github.com/KbaYero/UTMStackDatasourcesMISP/models"

// cleanData converts an object obtained from the MISP api, into an object to expose in the feeds url
func cleanData(data models.EventsRestSearch) models.CleanedEventsBody {
	var cleanedData models.CleanedEventsBody
	cleanedData.Events.Analysis = data.Events.Analysis
	cleanedData.Events.Date = data.Events.Date
	cleanedData.Events.ExtendsUUID = data.Events.ExtendsUUID
	cleanedData.Events.Info = data.Events.Info
	cleanedData.Events.PublishTimestamp = data.Events.PublishTimestamp
	cleanedData.Events.Published = data.Events.Published
	cleanedData.Events.ThreatLevelID = data.Events.ThreatLevelID
	cleanedData.Events.Timestamp = data.Events.Timestamp
	cleanedData.Events.UUID = data.Events.UUID
	cleanedData.Events.Orgc.Name = data.Events.Orgc.Name
	cleanedData.Events.Orgc.UUID = data.Events.Orgc.UUID
	for _, tag := range data.Events.Tag {
		var newTag models.TagCleaned
		newTag.Colour = tag.Colour
		newTag.Name = tag.Name
		cleanedData.Events.Tag = append(cleanedData.Events.Tag, newTag)
	}
	for _, attr := range data.Events.Attribute {
		var newAttribute models.AttributeCleaned
		newAttribute.Category = attr.Category
		newAttribute.Comment = attr.Comment
		newAttribute.Deleted = attr.Deleted
		newAttribute.DisableCorrelation = attr.DisableCorrelation
		newAttribute.Timestamp = attr.Timestamp
		newAttribute.ToIDS = attr.ToIDS
		newAttribute.Type = attr.Type
		newAttribute.UUID = attr.UUID
		newAttribute.Value = attr.Value
		cleanedData.Events.Attribute = append(cleanedData.Events.Attribute, newAttribute)
	}
	return cleanedData
}
