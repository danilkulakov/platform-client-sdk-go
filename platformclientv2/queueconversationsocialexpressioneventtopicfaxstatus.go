package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationsocialexpressioneventtopicfaxstatus
type Queueconversationsocialexpressioneventtopicfaxstatus struct { 
	// Direction
	Direction *string `json:"direction,omitempty"`


	// ExpectedPages
	ExpectedPages *int `json:"expectedPages,omitempty"`


	// ActivePage
	ActivePage *int `json:"activePage,omitempty"`


	// LinesTransmitted
	LinesTransmitted *int `json:"linesTransmitted,omitempty"`


	// BytesTransmitted
	BytesTransmitted *int `json:"bytesTransmitted,omitempty"`


	// BaudRate
	BaudRate *int `json:"baudRate,omitempty"`


	// PageErrors
	PageErrors *int `json:"pageErrors,omitempty"`


	// LineErrors
	LineErrors *int `json:"lineErrors,omitempty"`

}

func (o *Queueconversationsocialexpressioneventtopicfaxstatus) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationsocialexpressioneventtopicfaxstatus
	
	return json.Marshal(&struct { 
		Direction *string `json:"direction,omitempty"`
		
		ExpectedPages *int `json:"expectedPages,omitempty"`
		
		ActivePage *int `json:"activePage,omitempty"`
		
		LinesTransmitted *int `json:"linesTransmitted,omitempty"`
		
		BytesTransmitted *int `json:"bytesTransmitted,omitempty"`
		
		BaudRate *int `json:"baudRate,omitempty"`
		
		PageErrors *int `json:"pageErrors,omitempty"`
		
		LineErrors *int `json:"lineErrors,omitempty"`
		*Alias
	}{ 
		Direction: o.Direction,
		
		ExpectedPages: o.ExpectedPages,
		
		ActivePage: o.ActivePage,
		
		LinesTransmitted: o.LinesTransmitted,
		
		BytesTransmitted: o.BytesTransmitted,
		
		BaudRate: o.BaudRate,
		
		PageErrors: o.PageErrors,
		
		LineErrors: o.LineErrors,
		Alias:    (*Alias)(o),
	})
}

func (o *Queueconversationsocialexpressioneventtopicfaxstatus) UnmarshalJSON(b []byte) error {
	var QueueconversationsocialexpressioneventtopicfaxstatusMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationsocialexpressioneventtopicfaxstatusMap)
	if err != nil {
		return err
	}
	
	if Direction, ok := QueueconversationsocialexpressioneventtopicfaxstatusMap["direction"].(string); ok {
		o.Direction = &Direction
	}
	
	if ExpectedPages, ok := QueueconversationsocialexpressioneventtopicfaxstatusMap["expectedPages"].(float64); ok {
		ExpectedPagesInt := int(ExpectedPages)
		o.ExpectedPages = &ExpectedPagesInt
	}
	
	if ActivePage, ok := QueueconversationsocialexpressioneventtopicfaxstatusMap["activePage"].(float64); ok {
		ActivePageInt := int(ActivePage)
		o.ActivePage = &ActivePageInt
	}
	
	if LinesTransmitted, ok := QueueconversationsocialexpressioneventtopicfaxstatusMap["linesTransmitted"].(float64); ok {
		LinesTransmittedInt := int(LinesTransmitted)
		o.LinesTransmitted = &LinesTransmittedInt
	}
	
	if BytesTransmitted, ok := QueueconversationsocialexpressioneventtopicfaxstatusMap["bytesTransmitted"].(float64); ok {
		BytesTransmittedInt := int(BytesTransmitted)
		o.BytesTransmitted = &BytesTransmittedInt
	}
	
	if BaudRate, ok := QueueconversationsocialexpressioneventtopicfaxstatusMap["baudRate"].(float64); ok {
		BaudRateInt := int(BaudRate)
		o.BaudRate = &BaudRateInt
	}
	
	if PageErrors, ok := QueueconversationsocialexpressioneventtopicfaxstatusMap["pageErrors"].(float64); ok {
		PageErrorsInt := int(PageErrors)
		o.PageErrors = &PageErrorsInt
	}
	
	if LineErrors, ok := QueueconversationsocialexpressioneventtopicfaxstatusMap["lineErrors"].(float64); ok {
		LineErrorsInt := int(LineErrors)
		o.LineErrors = &LineErrorsInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationsocialexpressioneventtopicfaxstatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
