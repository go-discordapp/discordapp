package discordapp

type PresenceUpdate struct {
	User         User         `json:"user"`
	Roles        []RoleID     `json:"roles"`
	Game         *Activity    `json:"activity"`
	GuildID      GuildID      `json:"guild_id"`
	Status       string       `json:"status"`
	Activities   []Activity   `json:"activities"`
	ClientStatus ClientStatus `json:"client_status"`
}

type ClientStatus struct {
	Desktop string `json:"desktop"`
	Mobile  string `json:"mobile"`
	Web     string `json:"web"`
}

type Activity struct {
	Name          string                `json:"name"`
	Type          int                   `json:"type"`
	URL           *string               `json:"url,omit_empty"`
	Timestamps    *[]ActivityTimestamps `json:"timestamps,omit_empty"`
	ApplicationID *ApplicationID        `json:"application_id,omit_empty"`
	Details       *string               `json:"details,omit_empty"`
	State         *string               `json:"state,omit_empty"`
	Party         *ActivityParty        `json:"party,omit_empty"`
	Assets        *ActivityAssets       `json:"assets,omit_empty"`
	Secrets       *ActivitySecrets      `json:"secrets,omit_empty"`
	Instance      *bool                 `json:"instance,omit_empty"`
	Flags         *int                  `json:"flags,omit_empty"`
}

type ActivityTimestamps struct {
	Start *int `json:"start,omit_empty"`
	End   *int `json:"end,omit_empty"`
}

type ActivityParty struct {
	ID   *string `json:"id,omit_empty"`
	Size *[2]int `json:"size,omit_empty"`
}

type ActivityAssets struct {
	LargeImage *string `json:"large_image,omit_empty"`
	LargeText  *string `json:"large_text,omit_empty"`
	SmallImage *string `json:"small_image,omit_empty"`
	SmallText  *string `json:"small_text,omit_empty"`
}

type ActivitySecrets struct {
	Join     *string `json:"join,omit_empty"`
	Spectate *string `json:"spectate,omit_empty"`
	Match    *string `json:"match,omit_empty"`
}
