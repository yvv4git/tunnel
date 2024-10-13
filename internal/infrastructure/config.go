package infrastructure

type (
	Config struct {
		Server Server `mapstructure:"Server"`
		Client Client `mapstructure:"Client"`
	}

	Server struct {
		ChannelType ChannelType `mapstructure:"ChannelType"`
		TCPConfig   TCPServer   `mapstructure:"TCPConfig"`
		DeviceTUN   DeviceTUN   `mapstructure:"DeviceTUN"`
	}

	Client struct {
		ChannelType ChannelType `mapstructure:"ChannelType"`
		TCPConfig   TCPClient   `mapstructure:"TCPConfig"`
		DeviceTUN   DeviceTUN   `mapstructure:"DeviceTUN"`
	}

	DeviceTUN struct {
		Host     string `mapstructure:"Host"`
		Route    string `mapstructure:"Route"`
		Platform string `mapstructure:"Platform"`
	}

	TCPServer struct {
		Host       string              `mapstructure:"Host"`
		Port       uint16              `mapstructure:"Port"`
		BufferSize uint16              `mapstructure:"BufferSize"`
		Encryption TCPServerEncryptoin `mapstructure:"Encryption"`
	}

	TCPServerEncryptoin struct {
		Enabled    bool   `mapstructure:"Enabled"`
		ServerCert string `mapstructure:"ServerCert"`
		ServerKey  string `mapstructure:"ServerKey"`
		CACert     string `mapstructure:"CACert"`
	}

	TCPClient struct {
		ServerHost string              `mapstructure:"ServerHost"`
		ServerPort uint16              `mapstructure:"ServerPort"`
		BufferSize uint16              `mapstructure:"BufferSize"`
		Encryption TCPClientEncryptoin `mapstructure:"Encryption"`
	}

	TCPClientEncryptoin struct {
		Enabled    bool   `mapstructure:"Enabled"`
		ClientCert string `mapstructure:"ClientCert"`
		ClientKey  string `mapstructure:"ClientKey"`
		CACert     string `mapstructure:"CACert"`
	}
)
