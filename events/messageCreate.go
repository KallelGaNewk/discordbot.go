package events

import "github.com/bwmarrin/discordgo"

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, getAvatar(m.Author))
	}
}

func getAvatar(user *discordgo.User) string {
	return "https://cdn.discordapp.com/avatars/" + user.ID + "/" + user.Avatar + ".png"
}
