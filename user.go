package discordapp

type User struct {
	ID            UserID  `json:"user_id"`
	Username      string  `json:"username"`
	Discriminator string  `json:"discriminator"`
	Avatar        *string `json:"avatar"`
	Bot           *bool   `json:"bot,omit_empty"`
	MFAEnabled    *bool   `json:"mfa_enabled,omit_empty"`
	Locale        *string `json:"locale,omit_empty"`
	Verified      *bool   `json:"verified,omit_empty"`
	Email         *string `json:"email,omit_empty"`
	Flags         *int    `json:"flags,omit_empty"`
	PermiumType   *int    `json:"premium_type,omit_empty"`
}
