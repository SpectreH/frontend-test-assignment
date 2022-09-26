package config

const (
	CONFIG_PATH string = "./config"
	CONFIG_NAME string = "config"
	CONFIG_TYPE string = "yml"
	UI_DIST     string = "../frontend/dist"
)

// FileConfig holds full struct of config.yml file
type FileConfig struct {
	Port string
}
