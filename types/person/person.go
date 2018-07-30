package person

import (
	"time"
)

type MelodyPerson struct {
	Tenant    int       `json:"tenant,omitempty"`
	ID        int       `json:"id,omitempty"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Email     string    `json:"email,omitempty"`
	AccountID int       `json:"account_id,omitempty"`
	Created   time.Time `json:"created,omitempty"`
	Updated   time.Time `json:"update,omitempty"`
}

const IndexName = "fullsearch_test"
const TypeName = "person"
const Mapping = `
{
	"settings":{
		"number_of_shards": 1,
		"number_of_replicas": 0
	},
	"mappings":{
		"person":{
			"properties":{
				"tenant":{
					"type":"integer"
				},
				"id":{
					"type":"integer"
				},
				"first_name":{
					"type":"keyword"
				},
				"last_name":{
					"type":"keyword"
				},
				"email":{
					"type":"keyword"
				},
				"account_id":{
					"type":"integer"
				},
				"created":{
					"type":"date"
				},
				"updated":{
					"type":"date"
				}
			}
		}
	}
}`
