package infrastructure

type (
	Config struct {
		Server Server `mapstructure:"Server"`
		Client Client `mapstructure:"Client"`
	}

	Server struct {
		Host      string    `mapstructure:"Host"`
		Port      uint16    `mapstructure:"Port"`
		DeviceTUN DeviceTUN `mapstructure:"DeviceTUNAddress"`
	}

	Client struct {
		ServerHost string    `mapstructure:"ServerHost"`
		ServerPort uint16    `mapstructure:"ServerPort"`
		DeviceTUN  DeviceTUN `mapstructure:"DeviceTUNAddress"`
	}

	DeviceTUN struct {
		Host  string `mapstructure:"Host"`
		Route string `mapstructure:"Route"`
	}
)
