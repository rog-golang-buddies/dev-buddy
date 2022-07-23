package config

// For reference, look at https://github.com/sethvargo/go-envconfig
// EnvironmentConfig contains the struct for reading all env vars
type EnvironmentConfig struct {
	DiscordToken string `env:"DISCORD_TOKEN"`
}
