package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Keyperformanceindicator
type Keyperformanceindicator struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The name of the Key Performance Indicator.
	Name *string `json:"name,omitempty"`

}

func (o *Keyperformanceindicator) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Keyperformanceindicator
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		Alias:    (*Alias)(o),
	})
}

func (o *Keyperformanceindicator) UnmarshalJSON(b []byte) error {
	var KeyperformanceindicatorMap map[string]interface{}
	err := json.Unmarshal(b, &KeyperformanceindicatorMap)
	if err != nil {
		return err
	}
	
	if Id, ok := KeyperformanceindicatorMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := KeyperformanceindicatorMap["name"].(string); ok {
		o.Name = &Name
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Keyperformanceindicator) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}