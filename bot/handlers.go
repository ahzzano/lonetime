package bot

import (
	"errors"

	"github.com/bwmarrin/discordgo"
)

var sessions []Session
var guildConfigs []GuildConfig

func findSessionByWaitingRoom(channelID string) string {
	for i := 0; i < len(sessions); i++ {
		if sessions[i].waitingRoomChannel == channelID {
			return sessions[i].waitingRoomChannel
		}
	}

	return ""
}

func getGuildConfig(guildID string) (GuildConfig, error) {
	for i := 0; i < len(guildConfigs); i++ {
		if guildConfigs[i].guildID == guildID {
			return guildConfigs[i], nil
		}
	}

	return GuildConfig{}, errors.New("Unable to find config")
}

func CreateVoiceSession(s *discordgo.Session, event *discordgo.VoiceStateUpdate) error {
	return nil
}

func CreateJoinRequest(s *discordgo.Session, event *discordgo.VoiceStateUpdate) error {
	return nil
}
