package models

type ResponseBodyGetEvents struct {
	Response []EventsRestSearch `json:"response"`
}

type EventsRestSearch struct {
	Events Event `json:"Event"`
}

type Event struct {
	ID                 string      `json:"id"`
	OrgID              string      `json:"org_id"`
	Distribution       string      `json:"distribution"`
	Info               string      `json:"info"`
	OrgcID             string      `json:"orgc_id"`
	UUID               string      `json:"uuid"`
	Date               string      `json:"date"`
	Published          bool        `json:"published"`
	Analysis           string      `json:"analysis"`
	AttributeCount     string      `json:"attribute_count"`
	Timestamp          string      `json:"timestamp"`
	SharingGroupID     string      `json:"sharing_group_id"`
	ProposalEmailLock  bool        `json:"proposal_email_lock"`
	Locked             bool        `json:"locked"`
	ThreatLevelID      string      `json:"threat_level_id"`
	PublishTimestamp   string      `json:"publish_timestamp"`
	SightingTimestamp  string      `json:"sighting_timestamp"`
	DisableCorrelation bool        `json:"disable_correlation"`
	ExtendsUUID        string      `json:"extends_uuid"`
	EventCreatorEmail  string      `json:"event_creator_email"`
	Feed               feed        `json:"Feed"`
	Org                org         `json:"Org"`
	Orgc               org         `json:"Orgc"`
	Attribute          []attribute `json:"Attribute"`
	ShadowAttribute    []attribute `json:"ShadowAttribute"`
	RelatedEvent       []event     `json:"RelatedEvent"`
	Galaxy             []galaxy    `json:"Galaxy"`
	Object             []object    `json:"Object"`
	EventReport        []event     `json:"EventReport"`
	Tag                []tag       `json:"Tag"`
	Event              event       `json:"Event"`
}

type feed struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Provider        string `json:"provider"`
	Url             string `json:"url"`
	Rules           string `json:"rules"`
	Enabled         bool   `json:"enabled"`
	Distribution    string `json:"distribution"`
	SharingGroupID  string `json:"sharing_group_id"`
	TagID           string `json:"tag_id"`
	Default         bool   `json:"default"`
	SourceFormat    string `json:"source_format"`
	FixedEvent      bool   `json:"fixed_event"`
	DeltaMerge      bool   `json:"delta_merge"`
	EventID         string `json:"event_id"`
	Publish         bool   `json:"publish"`
	OverrideIds     bool   `json:"override_ids"`
	Settings        string `json:"settings"`
	InputSource     string `json:"input_source"`
	DeleteLocalFile bool   `json:"delete_local_file"`
	LookupVisible   bool   `json:"lookup_visible"`
	Headers         string `json:"headers"`
	CachingEnabled  bool   `json:"caching_enabled"`
	ForceToIds      bool   `json:"force_to_ids"`
	OrgcID          string `json:"orgc_id"`
	CacheTimestamp  string `json:"cache_timestamp"`
}

type org struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	UUID string `json:"uuid"`
}

type attribute struct {
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

type galaxy struct {
	ID             string        `json:"id"`
	UUID           string        `json:"uuid"`
	Name           string        `json:"name"`
	Type           string        `json:"type"`
	Description    string        `json:"description"`
	Version        string        `json:"version"`
	Icon           string        `json:"icon"`
	Namespace      string        `json:"namespace"`
	KillChainOrder fraud_tactics `json:"kill_chain_order"`
}

type fraud_tactics struct {
	Tactics []string `json:"fraud-tactics"`
}

type object struct {
	ID              string      `json:"id"`
	Name            string      `json:"name"`
	MetaCategory    string      `json:"meta-category"`
	Description     string      `json:"description"`
	TemplateUUID    string      `json:"template_uuid"`
	TemplateVersion string      `json:"template_version"`
	EventID         string      `json:"event_id"`
	UUID            string      `json:"uuid"`
	Timestamp       string      `json:"timestamp"`
	Distribution    string      `json:"distribution"`
	SharingGroupID  string      `json:"sharing_group_id"`
	Comment         string      `json:"comment"`
	Deleted         bool        `json:"deleted"`
	FirstSeen       string      `json:"first_seen"`
	LastSeen        string      `json:"last_seen"`
	Attribute       []attribute `json:"Attribute"`
}

type event struct {
	ID                 string `json:"id"`
	OrgID              string `json:"org_id"`
	Distribution       string `json:"distribution"`
	Info               string `json:"info"`
	OrgcID             string `json:"orgc_id"`
	UUID               string `json:"uuid"`
	Date               string `json:"date"`
	Published          bool   `json:"published"`
	Analysis           string `json:"analysis"`
	AttributeCount     string `json:"attribute_count"`
	Timestamp          string `json:"timestamp"`
	SharingGroupID     string `json:"sharing_group_id"`
	ProposalEmailLock  bool   `json:"proposal_email_lock"`
	Locked             bool   `json:"locked"`
	ThreatLevelID      string `json:"threat_level_id"`
	PublishTimestamp   string `json:"publish_timestamp"`
	SightingTimestamp  string `json:"sighting_timestamp"`
	DisableCorrelation bool   `json:"disable_correlation"`
	ExtendsUUID        string `json:"extends_uuid"`
	EventCreatorEmail  string `json:"event_creator_email"`
}

type tag struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Colour         string `json:"colour"`
	Exportable     bool   `json:"exportable"`
	OrgID          string `json:"org_id"`
	UserID         string `json:"user_id"`
	HideTag        bool   `json:"hide_tag"`
	NumericalValue string `json:"numerical_value"`
	IsGalaxy       bool   `json:"is_galaxy"`
	IsCustomGalaxy bool   `json:"is_custom_galaxy"`
	Inherited      int    `json:"inherited"`
}

type RequestEventsRestSearch struct {
	Page         int    `json:"page"`
	Limit        int    `json:"limit"`
	ReturnFormat string `json:"returnFormat"`
}

type CleanedEventsBody struct {
	Events CleanedEvent `json:"Event"`
}

type CleanedEvent struct {
	Analysis         string             `json:"analysis"`
	Date             string             `json:"date"`
	ExtendsUUID      string             `json:"extends_uuid"`
	Info             string             `json:"info"`
	PublishTimestamp string             `json:"publish_timestamp"`
	Published        bool               `json:"published"`
	ThreatLevelID    string             `json:"threat_level_id"`
	Timestamp        string             `json:"timestamp"`
	UUID             string             `json:"uuid"`
	Orgc             OrgCleaned         `json:"Orgc"`
	Tag              []TagCleaned       `json:"Tag"`
	Attribute        []AttributeCleaned `json:"Attribute"`
}

type OrgCleaned struct {
	Name string `json:"name"`
	UUID string `json:"uuid"`
}

type TagCleaned struct {
	Colour string `json:"colour"`
	Name   string `json:"name"`
}

type AttributeCleaned struct {
	Category           string `json:"category"`
	Comment            string `json:"comment"`
	Deleted            bool   `json:"deleted"`
	DisableCorrelation bool   `json:"disable_correlation"`
	Timestamp          string `json:"timestamp"`
	ToIDS              bool   `json:"to_ids"`
	Type               string `json:"type"`
	UUID               string `json:"uuid"`
	Value              string `json:"value"`
}
