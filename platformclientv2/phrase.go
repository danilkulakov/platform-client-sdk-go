package platformclientv2
import (
	"encoding/json"
)

// Phrase
type Phrase struct { 
	// Text - The phrase text
	Text *string `json:"text,omitempty"`


	// Strictness - The phrase strictness, default value is null
	Strictness *string `json:"strictness,omitempty"`


	// Sentiment - The phrase sentiment, default value is Unspecified
	Sentiment *string `json:"sentiment,omitempty"`

}

// String returns a JSON representation of the model
func (o *Phrase) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
