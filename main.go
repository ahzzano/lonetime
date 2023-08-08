package main

import (
	"flag"
	"fmt"
	"lonetime/bot"
	"lonetime/utils"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	tokenPtr := flag.String("token", "", "bot token")

	fmt.Println(*tokenPtr)

	flag.Parse()
	if *tokenPtr == "" {
		fmt.Println("No token provided")
		return
	}

	discord, err := discordgo.New("Bot " + *tokenPtr)

	if err != nil {
		panic(err)
	}

	utils.CheckError(err)

	discord.AddHandler(bot.OnReady)
	discord.AddHandler(bot.OnChannelJoin)
	discord.AddHandler(bot.OnChannelLeave)
	discord.AddHandler(bot.OnMessageReact)
	discord.AddHandler(bot.OnCommand)

	bot.CreateCommand("setup", bot.SetupGuild).Register()

	err = discord.Open()

	utils.CheckError(err)
	sc := make(chan os.Signal, 1)

	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	discord.Close()
}
