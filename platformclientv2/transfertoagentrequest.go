package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Transfertoagentrequest
type Transfertoagentrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// TransferType
	TransferType *string `json:"transferType,omitempty"`

	// UserId - The id of the internal user.
	UserId *string `json:"userId,omitempty"`

	// UserName - The userName (like user’s email) of the internal user.
	UserName *string `json:"userName,omitempty"`

	// UserDisplayName - The name of the internal user.
	UserDisplayName *string `json:"userDisplayName,omitempty"`

	// Voicemail - If true, transfer to the voicemail inbox of the participant that is being replaced.
	Voicemail *bool `json:"voicemail,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Transfertoagentrequest) SetField(field string, fieldValue interface{}) {
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

func (o Transfertoagentrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Transfertoagentrequest
	
	return json.Marshal(&struct { 
		TransferType *string `json:"transferType,omitempty"`
		
		UserId *string `json:"userId,omitempty"`
		
		UserName *string `json:"userName,omitempty"`
		
		UserDisplayName *string `json:"userDisplayName,omitempty"`
		
		Voicemail *bool `json:"voicemail,omitempty"`
		Alias
	}{ 
		TransferType: o.TransferType,
		
		UserId: o.UserId,
		
		UserName: o.UserName,
		
		UserDisplayName: o.UserDisplayName,
		
		Voicemail: o.Voicemail,
		Alias:    (Alias)(o),
	})
}

func (o *Transfertoagentrequest) UnmarshalJSON(b []byte) error {
	var TransfertoagentrequestMap map[string]interface{}
	err := json.Unmarshal(b, &TransfertoagentrequestMap)
	if err != nil {
		return err
	}
	
	if TransferType, ok := TransfertoagentrequestMap["transferType"].(string); ok {
		o.TransferType = &TransferType
	}
    
	if UserId, ok := TransfertoagentrequestMap["userId"].(string); ok {
		o.UserId = &UserId
	}
    
	if UserName, ok := TransfertoagentrequestMap["userName"].(string); ok {
		o.UserName = &UserName
	}
    
	if UserDisplayName, ok := TransfertoagentrequestMap["userDisplayName"].(string); ok {
		o.UserDisplayName = &UserDisplayName
	}
    
	if Voicemail, ok := TransfertoagentrequestMap["voicemail"].(bool); ok {
		o.Voicemail = &Voicemail
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Transfertoagentrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
