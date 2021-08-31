package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Bulkrelationshipsrequest
type Bulkrelationshipsrequest struct { 
	// Entities
	Entities *[]Relationship `json:"entities,omitempty"`

}

func (o *Bulkrelationshipsrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Bulkrelationshipsrequest
	
	return json.Marshal(&struct { 
		Entities *[]Relationship `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Bulkrelationshipsrequest) UnmarshalJSON(b []byte) error {
	var BulkrelationshipsrequestMap map[string]interface{}
	err := json.Unmarshal(b, &BulkrelationshipsrequestMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := BulkrelationshipsrequestMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Bulkrelationshipsrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
