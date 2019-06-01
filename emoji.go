package discordapp

type Emoji struct {
	ID            EmojiID `json:"id"`
	Name          string  `json:"name"`
	Roles         *[]Role `json:"roles,omit_empty"`
	User          *User   `json:"user,omit_empty"`
	RequireColons *bool   `json:"require_colons,omit_empty"`
	Managed       *bool   `json:"managed,omit_empty"`
	Animated      *bool   `json:"animated,omit_empty"`
}
