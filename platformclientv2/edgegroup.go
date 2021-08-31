package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Edgegroup
type Edgegroup struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The name of the entity.
	Name *string `json:"name,omitempty"`


	// Description - The resource's description.
	Description *string `json:"description,omitempty"`


	// Version - The current version of the resource.
	Version *int `json:"version,omitempty"`


	// DateCreated - The date the resource was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - The date of the last modification to the resource. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// ModifiedBy - The ID of the user that last modified the resource.
	ModifiedBy *string `json:"modifiedBy,omitempty"`


	// CreatedBy - The ID of the user that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`


	// State - Indicates if the resource is active, inactive, or deleted.
	State *string `json:"state,omitempty"`


	// ModifiedByApp - The application that last modified the resource.
	ModifiedByApp *string `json:"modifiedByApp,omitempty"`


	// CreatedByApp - The application that created the resource.
	CreatedByApp *string `json:"createdByApp,omitempty"`


	// Managed - Is this edge group being managed remotely.
	Managed *bool `json:"managed,omitempty"`


	// Hybrid - Is this edge group hybrid.
	Hybrid *bool `json:"hybrid,omitempty"`


	// EdgeTrunkBaseAssignment - A trunk base settings assignment of trunkType \"EDGE\" to use for edge-to-edge communication.
	EdgeTrunkBaseAssignment *Trunkbaseassignment `json:"edgeTrunkBaseAssignment,omitempty"`


	// PhoneTrunkBases - Trunk base settings of trunkType \"PHONE\" to inherit to edge logical interface for phone communication.
	PhoneTrunkBases *[]Trunkbase `json:"phoneTrunkBases,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Edgegroup) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Edgegroup
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		ModifiedBy *string `json:"modifiedBy,omitempty"`
		
		CreatedBy *string `json:"createdBy,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		ModifiedByApp *string `json:"modifiedByApp,omitempty"`
		
		CreatedByApp *string `json:"createdByApp,omitempty"`
		
		Managed *bool `json:"managed,omitempty"`
		
		Hybrid *bool `json:"hybrid,omitempty"`
		
		EdgeTrunkBaseAssignment *Trunkbaseassignment `json:"edgeTrunkBaseAssignment,omitempty"`
		
		PhoneTrunkBases *[]Trunkbase `json:"phoneTrunkBases,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Description: o.Description,
		
		Version: o.Version,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		ModifiedBy: o.ModifiedBy,
		
		CreatedBy: o.CreatedBy,
		
		State: o.State,
		
		ModifiedByApp: o.ModifiedByApp,
		
		CreatedByApp: o.CreatedByApp,
		
		Managed: o.Managed,
		
		Hybrid: o.Hybrid,
		
		EdgeTrunkBaseAssignment: o.EdgeTrunkBaseAssignment,
		
		PhoneTrunkBases: o.PhoneTrunkBases,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Edgegroup) UnmarshalJSON(b []byte) error {
	var EdgegroupMap map[string]interface{}
	err := json.Unmarshal(b, &EdgegroupMap)
	if err != nil {
		return err
	}
	
	if Id, ok := EdgegroupMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := EdgegroupMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Description, ok := EdgegroupMap["description"].(string); ok {
		o.Description = &Description
	}
	
	if Version, ok := EdgegroupMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if dateCreatedString, ok := EdgegroupMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := EdgegroupMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if ModifiedBy, ok := EdgegroupMap["modifiedBy"].(string); ok {
		o.ModifiedBy = &ModifiedBy
	}
	
	if CreatedBy, ok := EdgegroupMap["createdBy"].(string); ok {
		o.CreatedBy = &CreatedBy
	}
	
	if State, ok := EdgegroupMap["state"].(string); ok {
		o.State = &State
	}
	
	if ModifiedByApp, ok := EdgegroupMap["modifiedByApp"].(string); ok {
		o.ModifiedByApp = &ModifiedByApp
	}
	
	if CreatedByApp, ok := EdgegroupMap["createdByApp"].(string); ok {
		o.CreatedByApp = &CreatedByApp
	}
	
	if Managed, ok := EdgegroupMap["managed"].(bool); ok {
		o.Managed = &Managed
	}
	
	if Hybrid, ok := EdgegroupMap["hybrid"].(bool); ok {
		o.Hybrid = &Hybrid
	}
	
	if EdgeTrunkBaseAssignment, ok := EdgegroupMap["edgeTrunkBaseAssignment"].(map[string]interface{}); ok {
		EdgeTrunkBaseAssignmentString, _ := json.Marshal(EdgeTrunkBaseAssignment)
		json.Unmarshal(EdgeTrunkBaseAssignmentString, &o.EdgeTrunkBaseAssignment)
	}
	
	if PhoneTrunkBases, ok := EdgegroupMap["phoneTrunkBases"].([]interface{}); ok {
		PhoneTrunkBasesString, _ := json.Marshal(PhoneTrunkBases)
		json.Unmarshal(PhoneTrunkBasesString, &o.PhoneTrunkBases)
	}
	
	if SelfUri, ok := EdgegroupMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Edgegroup) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
