// AUTOGENERATED - DO NOT EDIT

package tdlib

// ChatInviteLinkMembers Contains a list of chat members joined a chat by an invite link
type ChatInviteLinkMembers struct {
	tdCommon
	TotalCount int32                  `json:"total_count"` // Approximate total count of chat members found
	Members    []ChatInviteLinkMember `json:"members"`     // List of chat members, joined a chat by an invite link
}

// MessageType return the string telegram-type of ChatInviteLinkMembers
func (chatInviteLinkMembers *ChatInviteLinkMembers) MessageType() string {
	return "chatInviteLinkMembers"
}

// NewChatInviteLinkMembers creates a new ChatInviteLinkMembers
//
// @param totalCount Approximate total count of chat members found
// @param members List of chat members, joined a chat by an invite link
func NewChatInviteLinkMembers(totalCount int32, members []ChatInviteLinkMember) *ChatInviteLinkMembers {
	chatInviteLinkMembersTemp := ChatInviteLinkMembers{
		tdCommon:   tdCommon{Type: "chatInviteLinkMembers"},
		TotalCount: totalCount,
		Members:    members,
	}

	return &chatInviteLinkMembersTemp
}
