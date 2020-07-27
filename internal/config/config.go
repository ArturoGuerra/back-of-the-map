package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type (
	// Role holds config information for giving out roles based on the Watch array and give role ID
	Role struct {
		Watch []string `yaml:"watch"`
		Give  string   `yaml:"give"`
	}

	// ReactionRole stores information for reaction roles
	ReactionRole struct {
		Emoji string `yaml:"emoji"`
		Role  string `yaml:"role"`
	}

	// Config holds all configuration information
	Config struct {
		ReactionRoles map[string][]*ReactionRole `yaml:"reaction_roles"`
		Roles         []*Role                    `yaml:"roles"`
		GuildID       string                     `yaml:"guild_id"`
	}
)

// Load loads config from yaml
func Load(location string) (*Config, error) {
	data, err := ioutil.ReadFile(location)
	if err != nil {
		return nil, err
	}

	config := Config{}

	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
