package discordapp

import "context"

type GuildService service

func (s GuildService) GetGuild(id GuildID, ctx context.Context) (*Guild, *Response, error) {
	req, err := s.c.newRequest("GET", "guilds/"+string(id), nil)
	g := &Guild{}
	resp, err := s.c.do(ctx, req, g)
	return g, resp, err
}

type Guild struct {
	ID   GuildID `json:"id"`
	Name string  `json:"name"`
}
