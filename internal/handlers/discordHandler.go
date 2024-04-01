package handler

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log/slog"
)

type DiscordHanlder struct {
	BaseHandler
}

var (
	Token string
)

func init() {
	//LoadToken
}
func (g *DiscordHanlder) Init() {

	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}
	dg.AddHandler(messageCreate)
	slog.Info("Discord Handler Initiated")
}

func (g *DiscordHanlder) Start() {

	slog.Info("Discord Handler Started")
}

func (g DiscordHanlder) Stop() {

	slog.Info("Terminal Handler Stopped")

}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

}
