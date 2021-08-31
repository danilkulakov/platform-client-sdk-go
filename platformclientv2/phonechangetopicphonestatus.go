package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Phonechangetopicphonestatus
type Phonechangetopicphonestatus struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// OperationalStatus
	OperationalStatus *string `json:"operationalStatus,omitempty"`


	// Edge
	Edge *Phonechangetopicedgereference `json:"edge,omitempty"`


	// Provision
	Provision *Phonechangetopicprovisioninfo `json:"provision,omitempty"`


	// LineStatuses
	LineStatuses *[]Phonechangetopiclinestatus `json:"lineStatuses,omitempty"`

}

func (o *Phonechangetopicphonestatus) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Phonechangetopicphonestatus
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		OperationalStatus *string `json:"operationalStatus,omitempty"`
		
		Edge *Phonechangetopicedgereference `json:"edge,omitempty"`
		
		Provision *Phonechangetopicprovisioninfo `json:"provision,omitempty"`
		
		LineStatuses *[]Phonechangetopiclinestatus `json:"lineStatuses,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		OperationalStatus: o.OperationalStatus,
		
		Edge: o.Edge,
		
		Provision: o.Provision,
		
		LineStatuses: o.LineStatuses,
		Alias:    (*Alias)(o),
	})
}

func (o *Phonechangetopicphonestatus) UnmarshalJSON(b []byte) error {
	var PhonechangetopicphonestatusMap map[string]interface{}
	err := json.Unmarshal(b, &PhonechangetopicphonestatusMap)
	if err != nil {
		return err
	}
	
	if Id, ok := PhonechangetopicphonestatusMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if OperationalStatus, ok := PhonechangetopicphonestatusMap["operationalStatus"].(string); ok {
		o.OperationalStatus = &OperationalStatus
	}
	
	if Edge, ok := PhonechangetopicphonestatusMap["edge"].(map[string]interface{}); ok {
		EdgeString, _ := json.Marshal(Edge)
		json.Unmarshal(EdgeString, &o.Edge)
	}
	
	if Provision, ok := PhonechangetopicphonestatusMap["provision"].(map[string]interface{}); ok {
		ProvisionString, _ := json.Marshal(Provision)
		json.Unmarshal(ProvisionString, &o.Provision)
	}
	
	if LineStatuses, ok := PhonechangetopicphonestatusMap["lineStatuses"].([]interface{}); ok {
		LineStatusesString, _ := json.Marshal(LineStatuses)
		json.Unmarshal(LineStatusesString, &o.LineStatuses)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Phonechangetopicphonestatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
