package events

import "github.com/bwmarrin/discordgo"

func Ready(s *discordgo.Session, event *discordgo.Ready) {
	s.UpdateListeningStatus("golang")
}
