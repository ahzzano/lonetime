package main

import (
	"lonetime/bot"
	"lonetime/utils"

	"github.com/bwmarrin/discordgo"
)

func main() {
	discord, err := discordgo.New("Bot " + "")

	if err != nil {
		panic(err)
	}

	utils.CheckError(err)

	discord.AddHandler(bot.OnReady)
	discord.AddHandler(bot.OnChannelJoin)
	discord.AddHandler(bot.OnChannelLeave)
	discord.AddHandler(bot.OnMessageReact)

	discord.Open()

	discord.Close()
}
