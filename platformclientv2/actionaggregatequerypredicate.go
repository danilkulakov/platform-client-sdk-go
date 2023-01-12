package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Actionaggregatequerypredicate
type Actionaggregatequerypredicate struct { 
	// VarType - Optional type, can usually be inferred
	VarType *string `json:"type,omitempty"`


	// Dimension - Left hand side for dimension predicates
	Dimension *string `json:"dimension,omitempty"`


	// Operator - Optional operator, default is matches
	Operator *string `json:"operator,omitempty"`


	// Value - Right hand side for dimension predicates
	Value *string `json:"value,omitempty"`


	// VarRange - Right hand side for dimension predicates
	VarRange *Numericrange `json:"range,omitempty"`

}

func (o *Actionaggregatequerypredicate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Actionaggregatequerypredicate
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Dimension *string `json:"dimension,omitempty"`
		
		Operator *string `json:"operator,omitempty"`
		
		Value *string `json:"value,omitempty"`
		
		VarRange *Numericrange `json:"range,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Dimension: o.Dimension,
		
		Operator: o.Operator,
		
		Value: o.Value,
		
		VarRange: o.VarRange,
		Alias:    (*Alias)(o),
	})
}

func (o *Actionaggregatequerypredicate) UnmarshalJSON(b []byte) error {
	var ActionaggregatequerypredicateMap map[string]interface{}
	err := json.Unmarshal(b, &ActionaggregatequerypredicateMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := ActionaggregatequerypredicateMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Dimension, ok := ActionaggregatequerypredicateMap["dimension"].(string); ok {
		o.Dimension = &Dimension
	}
    
	if Operator, ok := ActionaggregatequerypredicateMap["operator"].(string); ok {
		o.Operator = &Operator
	}
    
	if Value, ok := ActionaggregatequerypredicateMap["value"].(string); ok {
		o.Value = &Value
	}
    
	if VarRange, ok := ActionaggregatequerypredicateMap["range"].(map[string]interface{}); ok {
		VarRangeString, _ := json.Marshal(VarRange)
		json.Unmarshal(VarRangeString, &o.VarRange)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Actionaggregatequerypredicate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}