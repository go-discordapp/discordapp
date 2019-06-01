package discordapp

import "time"

type Channel struct {
	ID                   ChannelID      `json:"id"`
	Type                 int            `json:"type"`
	GuildID              *GuildID       `json:"guild_id,omit_empty"`
	Position             *int           `json:"position,omit_empty"`
	PermissionOverwrites *Overwrite     `json:"permission_overwrites,omit_empty"`
	Name                 *string        `json:"name,omit_empty"`
	Topic                *string        `json:"topic,omit_empty"`
	NSFW                 *bool          `json:"nsfw,omit_empty"`
	LastMessageID        *MessageID     `json:"last_message_id,omit_empty"`
	Bitrate              *int           `json:"bitrate,omit_empty"`
	UserLimit            *int           `json:"user_limit,omit_empty"`
	RateLimitPerUser     *int           `json:"rate_limit_per_user,omit_empty"`
	Recipients           *[]User        `json:"recipients,omit_empty"`
	Icon                 *string        `json:"icon,omit_empty"`
	OwnerID              *UserID        `json:"owner_id,omit_empty"`
	ApplicationID        *ApplicationID `json:"application_id,omit_empty"`
	ParentID             *ChannelID     `json:"parent_id,omit_empty"`
	LastPinTimestamp     *time.Time     `json:"last_pin,omit_empty"`
}

type Overwrite struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Allow int    `json:"allow"`
	Deny  int    `json:"deny"`
}
