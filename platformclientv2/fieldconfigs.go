package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Fieldconfigs
type Fieldconfigs struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Org
	Org *Fieldconfig `json:"org,omitempty"`

	// Person
	Person *Fieldconfig `json:"person,omitempty"`

	// Group
	Group *Fieldconfig `json:"group,omitempty"`

	// ExternalContact
	ExternalContact *Fieldconfig `json:"externalContact,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Fieldconfigs) SetField(field string, fieldValue interface{}) {
	// Get Value object for field
	target := reflect.ValueOf(o)
	targetField := reflect.Indirect(target).FieldByName(field)

	// Set value
	if fieldValue != nil {
		targetField.Set(reflect.ValueOf(fieldValue))
	} else {
		// Must create a new Value (creates **type) then get its element (*type), which will be nil pointer of the appropriate type
		x := reflect.Indirect(reflect.New(targetField.Type()))
		targetField.Set(x)
	}

	// Add field to set field names list
	if o.SetFieldNames == nil {
		o.SetFieldNames = make(map[string]bool)
	}
	o.SetFieldNames[field] = true
}

func (o Fieldconfigs) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if fieldValue == nil || reflect.ValueOf(fieldValue).IsNil()  {
				// Do nothing. Just catching this case to avoid trying to custom serialize a nil value
			} else if contains(dateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%fZ")
			} else if contains(localDateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%f")
			} else if contains(dateFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%d")
			}

			// Assign value to field using JSON tag name
			newObj[getFieldName(reflect.TypeOf(&o), fieldName)] = fieldValue
		}

		// Marshal and return dynamically constructed interface
		return json.Marshal(newObj)
	}

	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Fieldconfigs
	
	return json.Marshal(&struct { 
		Org *Fieldconfig `json:"org,omitempty"`
		
		Person *Fieldconfig `json:"person,omitempty"`
		
		Group *Fieldconfig `json:"group,omitempty"`
		
		ExternalContact *Fieldconfig `json:"externalContact,omitempty"`
		Alias
	}{ 
		Org: o.Org,
		
		Person: o.Person,
		
		Group: o.Group,
		
		ExternalContact: o.ExternalContact,
		Alias:    (Alias)(o),
	})
}

func (o *Fieldconfigs) UnmarshalJSON(b []byte) error {
	var FieldconfigsMap map[string]interface{}
	err := json.Unmarshal(b, &FieldconfigsMap)
	if err != nil {
		return err
	}
	
	if Org, ok := FieldconfigsMap["org"].(map[string]interface{}); ok {
		OrgString, _ := json.Marshal(Org)
		json.Unmarshal(OrgString, &o.Org)
	}
	
	if Person, ok := FieldconfigsMap["person"].(map[string]interface{}); ok {
		PersonString, _ := json.Marshal(Person)
		json.Unmarshal(PersonString, &o.Person)
	}
	
	if Group, ok := FieldconfigsMap["group"].(map[string]interface{}); ok {
		GroupString, _ := json.Marshal(Group)
		json.Unmarshal(GroupString, &o.Group)
	}
	
	if ExternalContact, ok := FieldconfigsMap["externalContact"].(map[string]interface{}); ok {
		ExternalContactString, _ := json.Marshal(ExternalContact)
		json.Unmarshal(ExternalContactString, &o.ExternalContact)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Fieldconfigs) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
