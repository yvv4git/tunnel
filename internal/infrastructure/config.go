package infrastructure

type (
	Config struct {
		Server Server `mapstructure:"Server"`
		Client Client `mapstructure:"Client"`
	}

	Server struct {
		Host string `mapstructure:"Host"`
		Port int    `mapstructure:"Port"`
	}

	Client struct {
		ServerHost string `mapstructure:"ServerHost"`
		ServerPort int    `mapstructure:"ServerPort"`
	}
)
