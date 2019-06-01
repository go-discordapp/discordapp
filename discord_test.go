package discordapp_test

import (
	"os"
	"testing"

	"github.com/go-discordapp/discordapp"
)

var (
	c       *discordapp.Client
	guildID = discordapp.GuildID("253947570953519104")
)

func setup(t *testing.T) {
	token := os.Getenv("TOKEN")
	token = "NTg0NDA0OTE1NzUzODQ0Nzg0.XPKbkA.NyHEXgtY5eGuj_FXYAq4YZO6QEM"
	if token == "" {
		t.Log("You need to set TOKEN enviroment value!!")
		t.FailNow()
	}
	c = discordapp.NewBotClient(token, nil)
}
