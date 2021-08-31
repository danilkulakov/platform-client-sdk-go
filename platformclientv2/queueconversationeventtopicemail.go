package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationeventtopicemail
type Queueconversationeventtopicemail struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// State
	State *string `json:"state,omitempty"`


	// Held
	Held *bool `json:"held,omitempty"`


	// AutoGenerated
	AutoGenerated *bool `json:"autoGenerated,omitempty"`


	// Subject
	Subject *string `json:"subject,omitempty"`


	// Provider
	Provider *string `json:"provider,omitempty"`


	// ScriptId
	ScriptId *string `json:"scriptId,omitempty"`


	// PeerId
	PeerId *string `json:"peerId,omitempty"`


	// MessagesSent
	MessagesSent *int `json:"messagesSent,omitempty"`


	// ErrorInfo
	ErrorInfo *Queueconversationeventtopicerrordetails `json:"errorInfo,omitempty"`


	// DisconnectType
	DisconnectType *string `json:"disconnectType,omitempty"`


	// StartHoldTime
	StartHoldTime *time.Time `json:"startHoldTime,omitempty"`


	// ConnectedTime
	ConnectedTime *time.Time `json:"connectedTime,omitempty"`


	// DisconnectedTime
	DisconnectedTime *time.Time `json:"disconnectedTime,omitempty"`


	// MessageId
	MessageId *string `json:"messageId,omitempty"`


	// Direction
	Direction *string `json:"direction,omitempty"`


	// DraftAttachments
	DraftAttachments *[]Queueconversationeventtopicattachment `json:"draftAttachments,omitempty"`


	// Spam
	Spam *bool `json:"spam,omitempty"`


	// Wrapup
	Wrapup *Queueconversationeventtopicwrapup `json:"wrapup,omitempty"`


	// AfterCallWork
	AfterCallWork *Queueconversationeventtopicaftercallwork `json:"afterCallWork,omitempty"`


	// AfterCallWorkRequired
	AfterCallWorkRequired *bool `json:"afterCallWorkRequired,omitempty"`


	// AdditionalProperties
	AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`

}

func (o *Queueconversationeventtopicemail) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationeventtopicemail
	
	StartHoldTime := new(string)
	if o.StartHoldTime != nil {
		
		*StartHoldTime = timeutil.Strftime(o.StartHoldTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartHoldTime = nil
	}
	
	ConnectedTime := new(string)
	if o.ConnectedTime != nil {
		
		*ConnectedTime = timeutil.Strftime(o.ConnectedTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ConnectedTime = nil
	}
	
	DisconnectedTime := new(string)
	if o.DisconnectedTime != nil {
		
		*DisconnectedTime = timeutil.Strftime(o.DisconnectedTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DisconnectedTime = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		Held *bool `json:"held,omitempty"`
		
		AutoGenerated *bool `json:"autoGenerated,omitempty"`
		
		Subject *string `json:"subject,omitempty"`
		
		Provider *string `json:"provider,omitempty"`
		
		ScriptId *string `json:"scriptId,omitempty"`
		
		PeerId *string `json:"peerId,omitempty"`
		
		MessagesSent *int `json:"messagesSent,omitempty"`
		
		ErrorInfo *Queueconversationeventtopicerrordetails `json:"errorInfo,omitempty"`
		
		DisconnectType *string `json:"disconnectType,omitempty"`
		
		StartHoldTime *string `json:"startHoldTime,omitempty"`
		
		ConnectedTime *string `json:"connectedTime,omitempty"`
		
		DisconnectedTime *string `json:"disconnectedTime,omitempty"`
		
		MessageId *string `json:"messageId,omitempty"`
		
		Direction *string `json:"direction,omitempty"`
		
		DraftAttachments *[]Queueconversationeventtopicattachment `json:"draftAttachments,omitempty"`
		
		Spam *bool `json:"spam,omitempty"`
		
		Wrapup *Queueconversationeventtopicwrapup `json:"wrapup,omitempty"`
		
		AfterCallWork *Queueconversationeventtopicaftercallwork `json:"afterCallWork,omitempty"`
		
		AfterCallWorkRequired *bool `json:"afterCallWorkRequired,omitempty"`
		
		AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		State: o.State,
		
		Held: o.Held,
		
		AutoGenerated: o.AutoGenerated,
		
		Subject: o.Subject,
		
		Provider: o.Provider,
		
		ScriptId: o.ScriptId,
		
		PeerId: o.PeerId,
		
		MessagesSent: o.MessagesSent,
		
		ErrorInfo: o.ErrorInfo,
		
		DisconnectType: o.DisconnectType,
		
		StartHoldTime: StartHoldTime,
		
		ConnectedTime: ConnectedTime,
		
		DisconnectedTime: DisconnectedTime,
		
		MessageId: o.MessageId,
		
		Direction: o.Direction,
		
		DraftAttachments: o.DraftAttachments,
		
		Spam: o.Spam,
		
		Wrapup: o.Wrapup,
		
		AfterCallWork: o.AfterCallWork,
		
		AfterCallWorkRequired: o.AfterCallWorkRequired,
		
		AdditionalProperties: o.AdditionalProperties,
		Alias:    (*Alias)(o),
	})
}

func (o *Queueconversationeventtopicemail) UnmarshalJSON(b []byte) error {
	var QueueconversationeventtopicemailMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationeventtopicemailMap)
	if err != nil {
		return err
	}
	
	if Id, ok := QueueconversationeventtopicemailMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if State, ok := QueueconversationeventtopicemailMap["state"].(string); ok {
		o.State = &State
	}
	
	if Held, ok := QueueconversationeventtopicemailMap["held"].(bool); ok {
		o.Held = &Held
	}
	
	if AutoGenerated, ok := QueueconversationeventtopicemailMap["autoGenerated"].(bool); ok {
		o.AutoGenerated = &AutoGenerated
	}
	
	if Subject, ok := QueueconversationeventtopicemailMap["subject"].(string); ok {
		o.Subject = &Subject
	}
	
	if Provider, ok := QueueconversationeventtopicemailMap["provider"].(string); ok {
		o.Provider = &Provider
	}
	
	if ScriptId, ok := QueueconversationeventtopicemailMap["scriptId"].(string); ok {
		o.ScriptId = &ScriptId
	}
	
	if PeerId, ok := QueueconversationeventtopicemailMap["peerId"].(string); ok {
		o.PeerId = &PeerId
	}
	
	if MessagesSent, ok := QueueconversationeventtopicemailMap["messagesSent"].(float64); ok {
		MessagesSentInt := int(MessagesSent)
		o.MessagesSent = &MessagesSentInt
	}
	
	if ErrorInfo, ok := QueueconversationeventtopicemailMap["errorInfo"].(map[string]interface{}); ok {
		ErrorInfoString, _ := json.Marshal(ErrorInfo)
		json.Unmarshal(ErrorInfoString, &o.ErrorInfo)
	}
	
	if DisconnectType, ok := QueueconversationeventtopicemailMap["disconnectType"].(string); ok {
		o.DisconnectType = &DisconnectType
	}
	
	if startHoldTimeString, ok := QueueconversationeventtopicemailMap["startHoldTime"].(string); ok {
		StartHoldTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startHoldTimeString)
		o.StartHoldTime = &StartHoldTime
	}
	
	if connectedTimeString, ok := QueueconversationeventtopicemailMap["connectedTime"].(string); ok {
		ConnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", connectedTimeString)
		o.ConnectedTime = &ConnectedTime
	}
	
	if disconnectedTimeString, ok := QueueconversationeventtopicemailMap["disconnectedTime"].(string); ok {
		DisconnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", disconnectedTimeString)
		o.DisconnectedTime = &DisconnectedTime
	}
	
	if MessageId, ok := QueueconversationeventtopicemailMap["messageId"].(string); ok {
		o.MessageId = &MessageId
	}
	
	if Direction, ok := QueueconversationeventtopicemailMap["direction"].(string); ok {
		o.Direction = &Direction
	}
	
	if DraftAttachments, ok := QueueconversationeventtopicemailMap["draftAttachments"].([]interface{}); ok {
		DraftAttachmentsString, _ := json.Marshal(DraftAttachments)
		json.Unmarshal(DraftAttachmentsString, &o.DraftAttachments)
	}
	
	if Spam, ok := QueueconversationeventtopicemailMap["spam"].(bool); ok {
		o.Spam = &Spam
	}
	
	if Wrapup, ok := QueueconversationeventtopicemailMap["wrapup"].(map[string]interface{}); ok {
		WrapupString, _ := json.Marshal(Wrapup)
		json.Unmarshal(WrapupString, &o.Wrapup)
	}
	
	if AfterCallWork, ok := QueueconversationeventtopicemailMap["afterCallWork"].(map[string]interface{}); ok {
		AfterCallWorkString, _ := json.Marshal(AfterCallWork)
		json.Unmarshal(AfterCallWorkString, &o.AfterCallWork)
	}
	
	if AfterCallWorkRequired, ok := QueueconversationeventtopicemailMap["afterCallWorkRequired"].(bool); ok {
		o.AfterCallWorkRequired = &AfterCallWorkRequired
	}
	
	if AdditionalProperties, ok := QueueconversationeventtopicemailMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationeventtopicemail) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
