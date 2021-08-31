package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Architectpromptnotificationarchitectoperation
type Architectpromptnotificationarchitectoperation struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Complete
	Complete *bool `json:"complete,omitempty"`


	// User
	User *Architectpromptnotificationuser `json:"user,omitempty"`


	// Client
	Client *Architectpromptnotificationclient `json:"client,omitempty"`


	// ActionName
	ActionName *string `json:"actionName,omitempty"`


	// ActionStatus
	ActionStatus *string `json:"actionStatus,omitempty"`


	// ErrorMessage
	ErrorMessage *string `json:"errorMessage,omitempty"`


	// ErrorCode
	ErrorCode *string `json:"errorCode,omitempty"`


	// ErrorMessageParams
	ErrorMessageParams *Architectpromptnotificationerrormessageparams `json:"errorMessageParams,omitempty"`


	// ErrorDetails
	ErrorDetails *[]Architectpromptnotificationerrordetail `json:"errorDetails,omitempty"`

}

func (o *Architectpromptnotificationarchitectoperation) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Architectpromptnotificationarchitectoperation
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Complete *bool `json:"complete,omitempty"`
		
		User *Architectpromptnotificationuser `json:"user,omitempty"`
		
		Client *Architectpromptnotificationclient `json:"client,omitempty"`
		
		ActionName *string `json:"actionName,omitempty"`
		
		ActionStatus *string `json:"actionStatus,omitempty"`
		
		ErrorMessage *string `json:"errorMessage,omitempty"`
		
		ErrorCode *string `json:"errorCode,omitempty"`
		
		ErrorMessageParams *Architectpromptnotificationerrormessageparams `json:"errorMessageParams,omitempty"`
		
		ErrorDetails *[]Architectpromptnotificationerrordetail `json:"errorDetails,omitempty"`
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

func (o *Architectpromptnotificationarchitectoperation) UnmarshalJSON(b []byte) error {
	var ArchitectpromptnotificationarchitectoperationMap map[string]interface{}
	err := json.Unmarshal(b, &ArchitectpromptnotificationarchitectoperationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ArchitectpromptnotificationarchitectoperationMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Complete, ok := ArchitectpromptnotificationarchitectoperationMap["complete"].(bool); ok {
		o.Complete = &Complete
	}
	
	if User, ok := ArchitectpromptnotificationarchitectoperationMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if Client, ok := ArchitectpromptnotificationarchitectoperationMap["client"].(map[string]interface{}); ok {
		ClientString, _ := json.Marshal(Client)
		json.Unmarshal(ClientString, &o.Client)
	}
	
	if ActionName, ok := ArchitectpromptnotificationarchitectoperationMap["actionName"].(string); ok {
		o.ActionName = &ActionName
	}
	
	if ActionStatus, ok := ArchitectpromptnotificationarchitectoperationMap["actionStatus"].(string); ok {
		o.ActionStatus = &ActionStatus
	}
	
	if ErrorMessage, ok := ArchitectpromptnotificationarchitectoperationMap["errorMessage"].(string); ok {
		o.ErrorMessage = &ErrorMessage
	}
	
	if ErrorCode, ok := ArchitectpromptnotificationarchitectoperationMap["errorCode"].(string); ok {
		o.ErrorCode = &ErrorCode
	}
	
	if ErrorMessageParams, ok := ArchitectpromptnotificationarchitectoperationMap["errorMessageParams"].(map[string]interface{}); ok {
		ErrorMessageParamsString, _ := json.Marshal(ErrorMessageParams)
		json.Unmarshal(ErrorMessageParamsString, &o.ErrorMessageParams)
	}
	
	if ErrorDetails, ok := ArchitectpromptnotificationarchitectoperationMap["errorDetails"].([]interface{}); ok {
		ErrorDetailsString, _ := json.Marshal(ErrorDetails)
		json.Unmarshal(ErrorDetailsString, &o.ErrorDetails)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Architectpromptnotificationarchitectoperation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
