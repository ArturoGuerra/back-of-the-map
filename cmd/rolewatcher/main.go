package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/arturoguerra/go-logging"
	"github.com/arturoguerra/rolewatcher/internal/config"
	"github.com/arturoguerra/rolewatcher/internal/handlers"
	"github.com/bwmarrin/discordgo"
)

var cfgLocation = "/config/config.yaml"

func init() {
	if val := os.Getenv("CONFIG_LOCATION"); len(val) != 0 {
		cfgLocation = val
	}
}

func main() {
	log := logging.New()

	cfg, err := config.Load(cfgLocation)
	if err != nil {
		log.Error(err)
		return
	}

	token := os.Getenv("TOKEN")
	if len(token) == 0 {
		log.Error("Missing token")
		return
	}

	dgo, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Error(err)
		return
	}

	h, err := handlers.New(log, cfg)
	if err != nil {
		log.Error(err)
		return
	}

	dgo.AddHandler(h.OnReady)
	//dgo.AddHandler(h.ReactionRolesAdd)
	//dgo.AddHandler(h.ReactionRolesRemove)
	dgo.AddHandler(h.RoleWatcher)

	if err = dgo.Open(); err != nil {
		log.Error(err)
		return
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dgo.Close()
}
