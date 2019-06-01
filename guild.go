package discordapp

import (
	"context"
	"time"
)

type GuildService service

type Guild struct {
	ID                          GuildID          `json:"id"`
	Name                        string           `json:"name"`
	Icon                        *string          `json:"icon"`
	Splash                      *string          `json:"splash"`
	Owner                       *bool            `json:"owner,omit_empty"`
	OwnerID                     UserID           `json:"owner_id"`
	Permissions                 *int             `json:"permissions,omit_empty"`
	Region                      string           `json:"region"`
	AFKChannelID                *ChannelID       `json:"afk_channel_id"`
	AFKTimeout                  int              `json:"afk_timeout"`
	EmbedEnabled                *bool            `json:"embed_enabled,omit_empty"`
	EmbedChannelID              *ChannelID       `json:"embed_channel_id,omit_empty"`
	VerificationLevel           int              `json:"verification_level"`
	DefaultMessageNotifications int              `json:"default_message_notifications"`
	ExplicitContentFilter       int              `json:"explicit_content_filter"`
	Roles                       []Role           `json:"roles"`
	Emojis                      []Emoji          `json:"emojis"`
	Features                    []string         `json:"features"`
	MFALevel                    int              `json:"mfa_level"`
	ApplicationID               *ApplicationID   `json:"application_id"`
	WidgetEnabled               bool             `json:"widget_enabled"`
	WidgetChannelID             *WidgetChannelID `json:"widget_channel_id,omit_empty"`
	SystemChannelID             *SystemChannelID `json:"system_channel_id"`
	MaxPresences                int              `json:"max_presences"`
	MaxMembers                  int              `json:"max_members"`
	VanityURLCode               *string          `json:"vanity_url_code"`
	Banner                      *string          `json:"banner"`
}

type GuildCreated struct {
	ID                          GuildID          `json:"id"`
	Name                        string           `json:"name"`
	Icon                        *string          `json:"icon"`
	Splash                      *string          `json:"splash"`
	Owner                       *bool            `json:"owner,omit_empty"`
	OwnerID                     UserID           `json:"owner_id"`
	Permissions                 *int             `json:"permissions,omit_empty"`
	Region                      string           `json:"region"`
	AFKChannelID                *ChannelID       `json:"afk_channel_id"`
	AFKTimeout                  int              `json:"afk_timeout"`
	EmbedEnabled                *bool            `json:"embed_enabled,omit_empty"`
	EmbedChannelID              *ChannelID       `json:"embed_channel_id,omit_empty"`
	VerificationLevel           int              `json:"verification_level"`
	DefaultMessageNotifications int              `json:"default_message_notifications"`
	ExplicitContentFilter       int              `json:"explicit_content_filter"`
	Roles                       []Role           `json:"roles"`
	Emojis                      []Emoji          `json:"emojis"`
	Features                    []string         `json:"features"`
	MFALevel                    int              `json:"mfa_level"`
	ApplicationID               *ApplicationID   `json:"application_id"`
	WidgetEnabled               bool             `json:"widget_enabled"`
	WidgetChannelID             *WidgetChannelID `json:"widget_channel_id,omit_empty"`
	SystemChannelID             *SystemChannelID `json:"system_channel_id"`
	JoinedAt                    time.Time        `json:"joined_at"`
	Large                       bool             `json:"large"`
	Unavailable                 bool             `json:"unavailable"`
	MemberCount                 int              `json:"member_count"`
	VoiceStates                 []VoiceState     `json:"voice_states"`
	Members                     []GuildMember    `json:"members"`
	Channels                    []Channel        `json:"channels"`
	Presences                   []PresenceUpdate `json:"presences"`
	MaxPresences                int              `json:"max_presences"`
	MaxMembers                  int              `json:"max_members"`
	VanityURLCode               *string          `json:"vanity_url_code"`
	Banner                      *string          `json:"banner"`
}

type GuildMember struct {
	User     User      `json:"user"`
	Nick     *string   `json:"nick,omit_empty"`
	Roles    []RoleID  `json:"roles"`
	JoinedAt time.Time `json:"joined_at"`
	Deaf     bool      `json:"deaf"`
	Mute     bool      `json:"mute"`
}

func (s GuildService) GetGuild(id GuildID, ctx context.Context) (*Guild, *Response, error) {
	req, err := s.c.newRequest("GET", "guilds/"+string(id), nil)
	g := &Guild{}
	resp, err := s.c.do(ctx, req, g)
	return g, resp, err
}
