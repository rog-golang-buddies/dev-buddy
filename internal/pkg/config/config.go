package config

// For reference, look at https://github.com/sethvargo/go-envconfig
// EnvironmentConfig contains the struct for reading all env vars
type EnvironmentConfig struct {
	DiscordToken string `env:"DISCORD_TOKEN"`
	GithubPAT    string `env:"GITHUB_PAT"`
	OwnerName    string `env:"GITHUB_OWNER"`
}
