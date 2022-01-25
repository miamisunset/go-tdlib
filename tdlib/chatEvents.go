// AUTOGENERATED - DO NOT EDIT

package tdlib

// ChatEvents Contains a list of chat events
type ChatEvents struct {
	tdCommon
	Events []ChatEvent `json:"events"` // List of events
}

// MessageType return the string telegram-type of ChatEvents
func (chatEvents *ChatEvents) MessageType() string {
	return "chatEvents"
}

// NewChatEvents creates a new ChatEvents
//
// @param events List of events
func NewChatEvents(events []ChatEvent) *ChatEvents {
	chatEventsTemp := ChatEvents{
		tdCommon: tdCommon{Type: "chatEvents"},
		Events:   events,
	}

	return &chatEventsTemp
}
