package discordapp_test

import (
	"context"
	"testing"
)

func TestGuildService_GetGuild(t *testing.T) {
	setup(t)
	g, _, err := c.Guild.GetGuild(guildID, context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if g.Name != "go-discordapp" {
		t.Fatal("Guild name is not go-discordapp.")
	}
}
