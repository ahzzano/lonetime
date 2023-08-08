package bot

import "github.com/bwmarrin/discordgo"

func SetupGuild(s *discordgo.Session, event *discordgo.MessageCreate) error {
	s.ChannelMessageSend(event.ChannelID, "Hello There")
	return nil
}
