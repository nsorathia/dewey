package team

import (
	"github.com/olivere/elastic"
	"time"
)

type Team struct {
	Name             string                `json:"name,omitempty"`
	SubscriptionType string                `json:"subscription_type,omitempty"`
	LicenseLimit     int                   `json:"license_limit,omitempty"`
	Plan             string                `json:"plan,omitempty"`
	SoftLimit        bool                  `json:"soft_limit,omitempty"`
	Suggest          *elastic.SuggestField `json:"suggest_field,omitempty"`
	Created          time.Time             `json:"created,omitempty"`
	Updated          time.Time             `json:"update,omitempty"`
}

const IndexName = "team"
const TypeName = "team"

const Mapping = `
{
	"settings":{
		"number_of_shards": 1,
		"number_of_replicas": 0
	},
	"mappings":{
		"team":{
			"properties":{
				"name":{
					"type":"text"
				},
				"subscription_type":{
					"type":"keyword"
				},
				"license_limit":{
					"type":"integer"
				},
				"plan":{
					"type":"keyword"
				},
				"seat_type":{
					"type":"keyword"
				},
				"soft_limit":{
					"type":"integer"
				},
				"created":{
					"type":"date"
				},
				"updated":{
					"type":"date"
				},
				"suggest_field":{
					"type":"completion"
				}
			}
		}
	}
}`
