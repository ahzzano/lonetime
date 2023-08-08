package bot

import (
	"fmt"
	"lonetime/utils"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var onReadyHandlers []func(*discordgo.Session, *discordgo.Ready) error
var onChannelJoinHandlers []func(*discordgo.Session, *discordgo.VoiceStateUpdate) error
var onChannelLeaveHandlers []func(*discordgo.Session, *discordgo.VoiceStateUpdate) error
var onMessageReactHandlers []func(*discordgo.Session, *discordgo.MessageReactionAdd) error
var commands []Command

func AddReadyHandler(function func(*discordgo.Session, *discordgo.Ready) error) {
	onReadyHandlers = append(onReadyHandlers, function)
}

func AddChannelJoinHandler(function func(*discordgo.Session, *discordgo.VoiceStateUpdate) error) {
	onChannelJoinHandlers = append(onChannelJoinHandlers, function)
}

func AddChannelLeaveHandler(function func(*discordgo.Session, *discordgo.VoiceStateUpdate) error) {
	onChannelLeaveHandlers = append(onChannelLeaveHandlers, function)
}

func AddMessageReactHandlers(function func(*discordgo.Session, *discordgo.MessageReactionAdd) error) {
	onMessageReactHandlers = append(onMessageReactHandlers, function)
}

func OnReady(s *discordgo.Session, event *discordgo.Ready) {
	fmt.Println("Bot is Ready")

	for i := 0; i < len(onReadyHandlers); i++ {
		err := onReadyHandlers[i](s, event)

		utils.CheckError(err)
	}
}

func OnChannelJoin(s *discordgo.Session, event *discordgo.VoiceStateUpdate) {
	if event.BeforeUpdate.ChannelID == "" && event.ChannelID != "" {
		return
	}
	for i := 0; i < len(onChannelJoinHandlers); i++ {
		err := onChannelJoinHandlers[i](s, event)

		utils.CheckError(err)
	}
}

func OnChannelLeave(s *discordgo.Session, event *discordgo.VoiceStateUpdate) {
	if event.BeforeUpdate.ChannelID != "" && event.ChannelID == "" {
		return
	}

	for i := 0; i < len(onChannelLeaveHandlers); i++ {
		err := onChannelLeaveHandlers[i](s, event)

		utils.CheckError(err)
	}
}

func OnMessageReact(s *discordgo.Session, event *discordgo.MessageReactionAdd) {
	for i := 0; i < len(onMessageReactHandlers); i++ {
		err := onMessageReactHandlers[i](s, event)

		utils.CheckError(err)
	}
}

func OnCommand(s *discordgo.Session, event *discordgo.MessageCreate) {
	if !strings.HasPrefix(event.Content, prefix) {
		return
	}

	for i := 0; i < len(commands); i++ {
		if strings.Contains(event.Content, commands[i].command) {
			commands[i].action(s, event)
		}
	}
}
