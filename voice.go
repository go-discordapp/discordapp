package discordapp

type VoiceState struct {
	GuildID   *GuildID     `json:"guild_id,omit_empty"`
	ChannelID *ChannelID   `json:"channel_id"`
	UserID    User         `json:"user_id"`
	Member    *GuildMember `json:"guild_member,omit_empty"`
	SessionID string       `json:"session_id"`
	Deaf      bool         `json:"deaf"`
	Mute      bool         `json:"mute"`
	SelfDeaf  bool         `json:"self_deaf"`
	SelfMute  bool         `json:"self_mute"`
	Suppress  bool         `json:"suppress"`
}
