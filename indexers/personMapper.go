package indexers

import (
	"dewey/common"
	types "dewey/types/person"
)

//MapDocument parses the data map and returns a MelodyPerson type
func MapDocument(objectMap map[string]interface{}) (types.MelodyPerson, error) {

	context := objectMap["context"].(map[string]interface{})
	object := objectMap["object"].(map[string]interface{})

	tenant, error := common.GetInteger(context, "tenant")
	if error != nil {
		return types.MelodyPerson{}, error
	}

	id, error := common.GetInteger(object, "id")
	if error != nil {
		return types.MelodyPerson{}, error
	}

	firstname, error := common.GetString(object, "first_name")
	if error != nil {
		return types.MelodyPerson{}, error
	}

	lastname, error := common.GetString(object, "last_name")
	if error != nil {
		return types.MelodyPerson{}, error
	}

	email, error := common.GetString(object, "email_address")
	if error != nil {
		return types.MelodyPerson{}, error
	}

	//create a function to take stirng ("owner.id") and return string/Int, date
	owner := object["owner"].(map[string]interface{})
	accountID, error := common.GetInteger(owner, "id")
	if error != nil {
		return types.MelodyPerson{}, error
	}

	created, error := common.GetDate(object, "created_at")
	if error != nil {
		return types.MelodyPerson{}, error
	}

	updated, error := common.GetDate(object, "updated_at")
	if error != nil {
		return types.MelodyPerson{}, error
	}

	return types.MelodyPerson{
		ID:        id,
		Tenant:    tenant,
		FirstName: firstname,
		LastName:  lastname,
		Email:     email,
		AccountID: accountID,
		Created:   created,
		Updated:   updated,
	}, nil
}
