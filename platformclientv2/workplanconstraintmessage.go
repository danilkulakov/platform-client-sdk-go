package platformclientv2
import (
	"encoding/json"
)

// Workplanconstraintmessage
type Workplanconstraintmessage struct { 
	// VarType - Type of the work plan constraint in this message
	VarType *string `json:"type,omitempty"`


	// Arguments - Arguments of the message that provide information about the constraint that is being conflicted with, such as the value of the constraint
	Arguments *[]Workplanvalidationmessageargument `json:"arguments,omitempty"`

}

// String returns a JSON representation of the model
func (o *Workplanconstraintmessage) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}