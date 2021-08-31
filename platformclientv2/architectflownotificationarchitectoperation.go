package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Architectflownotificationarchitectoperation
type Architectflownotificationarchitectoperation struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Complete
	Complete *bool `json:"complete,omitempty"`


	// User
	User *Architectflownotificationuser `json:"user,omitempty"`


	// Client
	Client *Architectflownotificationclient `json:"client,omitempty"`


	// ActionName
	ActionName *string `json:"actionName,omitempty"`


	// ActionStatus
	ActionStatus *string `json:"actionStatus,omitempty"`


	// ErrorMessage
	ErrorMessage *string `json:"errorMessage,omitempty"`


	// ErrorCode
	ErrorCode *string `json:"errorCode,omitempty"`


	// ErrorMessageParams
	ErrorMessageParams *Architectflownotificationerrormessageparams `json:"errorMessageParams,omitempty"`


	// ErrorDetails
	ErrorDetails *[]Architectflownotificationerrordetail `json:"errorDetails,omitempty"`

}

func (o *Architectflownotificationarchitectoperation) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Architectflownotificationarchitectoperation
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Complete *bool `json:"complete,omitempty"`
		
		User *Architectflownotificationuser `json:"user,omitempty"`
		
		Client *Architectflownotificationclient `json:"client,omitempty"`
		
		ActionName *string `json:"actionName,omitempty"`
		
		ActionStatus *string `json:"actionStatus,omitempty"`
		
		ErrorMessage *string `json:"errorMessage,omitempty"`
		
		ErrorCode *string `json:"errorCode,omitempty"`
		
		ErrorMessageParams *Architectflownotificationerrormessageparams `json:"errorMessageParams,omitempty"`
		
		ErrorDetails *[]Architectflownotificationerrordetail `json:"errorDetails,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Complete: o.Complete,
		
		User: o.User,
		
		Client: o.Client,
		
		ActionName: o.ActionName,
		
		ActionStatus: o.ActionStatus,
		
		ErrorMessage: o.ErrorMessage,
		
		ErrorCode: o.ErrorCode,
		
		ErrorMessageParams: o.ErrorMessageParams,
		
		ErrorDetails: o.ErrorDetails,
		Alias:    (*Alias)(o),
	})
}

func (o *Architectflownotificationarchitectoperation) UnmarshalJSON(b []byte) error {
	var ArchitectflownotificationarchitectoperationMap map[string]interface{}
	err := json.Unmarshal(b, &ArchitectflownotificationarchitectoperationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ArchitectflownotificationarchitectoperationMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Complete, ok := ArchitectflownotificationarchitectoperationMap["complete"].(bool); ok {
		o.Complete = &Complete
	}
	
	if User, ok := ArchitectflownotificationarchitectoperationMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if Client, ok := ArchitectflownotificationarchitectoperationMap["client"].(map[string]interface{}); ok {
		ClientString, _ := json.Marshal(Client)
		json.Unmarshal(ClientString, &o.Client)
	}
	
	if ActionName, ok := ArchitectflownotificationarchitectoperationMap["actionName"].(string); ok {
		o.ActionName = &ActionName
	}
	
	if ActionStatus, ok := ArchitectflownotificationarchitectoperationMap["actionStatus"].(string); ok {
		o.ActionStatus = &ActionStatus
	}
	
	if ErrorMessage, ok := ArchitectflownotificationarchitectoperationMap["errorMessage"].(string); ok {
		o.ErrorMessage = &ErrorMessage
	}
	
	if ErrorCode, ok := ArchitectflownotificationarchitectoperationMap["errorCode"].(string); ok {
		o.ErrorCode = &ErrorCode
	}
	
	if ErrorMessageParams, ok := ArchitectflownotificationarchitectoperationMap["errorMessageParams"].(map[string]interface{}); ok {
		ErrorMessageParamsString, _ := json.Marshal(ErrorMessageParams)
		json.Unmarshal(ErrorMessageParamsString, &o.ErrorMessageParams)
	}
	
	if ErrorDetails, ok := ArchitectflownotificationarchitectoperationMap["errorDetails"].([]interface{}); ok {
		ErrorDetailsString, _ := json.Marshal(ErrorDetails)
		json.Unmarshal(ErrorDetailsString, &o.ErrorDetails)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Architectflownotificationarchitectoperation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
