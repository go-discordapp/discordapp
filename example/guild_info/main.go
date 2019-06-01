package main

import (
	"context"
	"os"

	"github.com/go-discordapp/discordapp"
	"github.com/k0kubun/pp"
)

func main() {
	token := os.Getenv("TOKEN")
	if token == "" {
		panic("Token is empty!!")
	}
	discord := discordapp.NewBotClient(token, nil)
	guild, _, err := discord.Guild.GetGuild("253947570953519104", context.Background())
	if err != nil {
		panic(err)
	}
	pp.Println(guild)
}
