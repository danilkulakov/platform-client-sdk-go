package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Architectflowoutcomenotificationarchitectoperation
type Architectflowoutcomenotificationarchitectoperation struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Complete
	Complete *bool `json:"complete,omitempty"`


	// User
	User *Architectflowoutcomenotificationuser `json:"user,omitempty"`


	// Client
	Client *Architectflowoutcomenotificationclient `json:"client,omitempty"`


	// ActionName
	ActionName *string `json:"actionName,omitempty"`


	// ActionStatus
	ActionStatus *string `json:"actionStatus,omitempty"`


	// ErrorMessage
	ErrorMessage *string `json:"errorMessage,omitempty"`


	// ErrorCode
	ErrorCode *string `json:"errorCode,omitempty"`


	// ErrorMessageParams
	ErrorMessageParams *Architectflowoutcomenotificationerrormessageparams `json:"errorMessageParams,omitempty"`


	// ErrorDetails
	ErrorDetails *[]Architectflowoutcomenotificationerrordetail `json:"errorDetails,omitempty"`

}

func (o *Architectflowoutcomenotificationarchitectoperation) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Architectflowoutcomenotificationarchitectoperation
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Complete *bool `json:"complete,omitempty"`
		
		User *Architectflowoutcomenotificationuser `json:"user,omitempty"`
		
		Client *Architectflowoutcomenotificationclient `json:"client,omitempty"`
		
		ActionName *string `json:"actionName,omitempty"`
		
		ActionStatus *string `json:"actionStatus,omitempty"`
		
		ErrorMessage *string `json:"errorMessage,omitempty"`
		
		ErrorCode *string `json:"errorCode,omitempty"`
		
		ErrorMessageParams *Architectflowoutcomenotificationerrormessageparams `json:"errorMessageParams,omitempty"`
		
		ErrorDetails *[]Architectflowoutcomenotificationerrordetail `json:"errorDetails,omitempty"`
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

func (o *Architectflowoutcomenotificationarchitectoperation) UnmarshalJSON(b []byte) error {
	var ArchitectflowoutcomenotificationarchitectoperationMap map[string]interface{}
	err := json.Unmarshal(b, &ArchitectflowoutcomenotificationarchitectoperationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ArchitectflowoutcomenotificationarchitectoperationMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Complete, ok := ArchitectflowoutcomenotificationarchitectoperationMap["complete"].(bool); ok {
		o.Complete = &Complete
	}
	
	if User, ok := ArchitectflowoutcomenotificationarchitectoperationMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if Client, ok := ArchitectflowoutcomenotificationarchitectoperationMap["client"].(map[string]interface{}); ok {
		ClientString, _ := json.Marshal(Client)
		json.Unmarshal(ClientString, &o.Client)
	}
	
	if ActionName, ok := ArchitectflowoutcomenotificationarchitectoperationMap["actionName"].(string); ok {
		o.ActionName = &ActionName
	}
	
	if ActionStatus, ok := ArchitectflowoutcomenotificationarchitectoperationMap["actionStatus"].(string); ok {
		o.ActionStatus = &ActionStatus
	}
	
	if ErrorMessage, ok := ArchitectflowoutcomenotificationarchitectoperationMap["errorMessage"].(string); ok {
		o.ErrorMessage = &ErrorMessage
	}
	
	if ErrorCode, ok := ArchitectflowoutcomenotificationarchitectoperationMap["errorCode"].(string); ok {
		o.ErrorCode = &ErrorCode
	}
	
	if ErrorMessageParams, ok := ArchitectflowoutcomenotificationarchitectoperationMap["errorMessageParams"].(map[string]interface{}); ok {
		ErrorMessageParamsString, _ := json.Marshal(ErrorMessageParams)
		json.Unmarshal(ErrorMessageParamsString, &o.ErrorMessageParams)
	}
	
	if ErrorDetails, ok := ArchitectflowoutcomenotificationarchitectoperationMap["errorDetails"].([]interface{}); ok {
		ErrorDetailsString, _ := json.Marshal(ErrorDetails)
		json.Unmarshal(ErrorDetailsString, &o.ErrorDetails)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Architectflowoutcomenotificationarchitectoperation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
