package bot

import "github.com/bwmarrin/discordgo"

var prefix string = "!"

// This struct contains all the information necessary for a lonetime session on Discord.
type Session struct {
	// Session Name
	name string

	// Guild ID of the session
	guild string

	// User ID of the owner
	owner string

	// Whether or not there are still users in the call
	active bool

	// Users who are allowed to enter the call
	acceptedUsers []string

	// Users that have joined and are in the waiting room
	joinRequests []WaitingRoomRequest

	// DiscordGo Related

	// Channel ID of the session
	mainChannel string

	// Channel ID of the waiting room
	waitingRoomChannel string

	// Channel ID of the text channel of the session
	textChannel string
}

type WaitingRoomRequest struct {
	// User ID of the user who wants to join the call
	user string

	// Whether the user is accepted or nope
	accepted bool
}

// How the guilds are setup
type GuildConfig struct {
	// The ID of the guild
	guildID string

	// ID of the generator channel
	generatorChannelID string

	// ID of the category channel
	categoryChannelID string
}

type Command struct {
	// The function that will be called
	action func(*discordgo.Session, *discordgo.MessageCreate) error
	// The command is the text in !command
	command string
}

func (c Command) Register() {
	commands = append(commands, c)
}

func CreateCommand(command string, action func(*discordgo.Session, *discordgo.MessageCreate) error) Command {
	newCommand := Command{
		command: command,
		action:  action,
	}

	return newCommand
}
