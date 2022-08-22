package main

import (
	"github.com/bwmarrin/discordgo"
)

func sendMessage(message string, messageType string, s *discordgo.Session, m *discordgo.MessageCreate) {
	channelId := m.ChannelID

	if messageType == "text" {
		s.ChannelMessageSend(channelId, message)
	}
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	prefix := getPrefix()

	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == prefix+"help" {
		sendMessage("uwu", "text", s, m)
	}
}
