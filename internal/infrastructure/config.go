package infrastructure

type (
	Config struct {
		DirectConnection DirectConnection `mapstructure:"DirectConnection"`
		SpeedTest        SpeedTest        `mapstructure:"SpeedTest"`
	}

	DirectConnection struct {
		Server Server `mapstructure:"Server"`
		Client Client `mapstructure:"Client"`
	}

	Server struct {
		ChannelType ChannelType `mapstructure:"ChannelType"`
		TCPConfig   TCPServer   `mapstructure:"TCPConfig"`
		UDPConfig   UDPServer   `mapstructure:"UDPConfig"`
		DeviceTUN   DeviceTUN   `mapstructure:"DeviceTUN"`
	}

	Client struct {
		ChannelType ChannelType `mapstructure:"ChannelType"`
		TCPConfig   TCPClient   `mapstructure:"TCPConfig"`
		UDPConfig   UDPClient   `mapstructure:"UDPConfig"`
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
		Metrics    MetricsWebServer    `mapstructure:"Metrics"`
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

	UDPServer struct {
		Host       string              `mapstructure:"Host"`
		Port       uint16              `mapstructure:"Port"`
		BufferSize uint16              `mapstructure:"BufferSize"`
		Encryption TCPServerEncryptoin `mapstructure:"Encryption"`
	}

	UDPClient struct {
		ServerHost string              `mapstructure:"ServerHost"`
		ServerPort uint16              `mapstructure:"ServerPort"`
		BufferSize uint16              `mapstructure:"BufferSize"`
		Encryption UDPClientEncryptoin `mapstructure:"Encryption"`
	}

	UDPClientEncryptoin struct {
		Enabled    bool   `mapstructure:"Enabled"`
		ClientCert string `mapstructure:"ClientCert"`
		ClientKey  string `mapstructure:"ClientKey"`
		CACert     string `mapstructure:"CACert"`
	}

	MetricsWebServer struct {
		Host string `mapstructure:"Host"`
		Port uint16 `mapstructure:"Port"`
	}

	SpeedTest struct {
		TCPServerSpeedTest TCPServerSpeedTest `mapstructure:"TCPServerSpeedTest"`
		TCPClientSpeedTest TCPClientSpeedTest `mapstructure:"TCPClientSpeedTest"`
	}

	TCPServerSpeedTest struct {
		Host       string              `mapstructure:"Host"`
		Port       uint16              `mapstructure:"Port"`
		BufferSize uint16              `mapstructure:"BufferSize"`
		Encryption TCPServerEncryptoin `mapstructure:"Encryption"`
	}

	TCPClientSpeedTest struct {
		ServerHost string              `mapstructure:"ServerHost"`
		ServerPort uint16              `mapstructure:"ServerPort"`
		BufferSize uint16              `mapstructure:"BufferSize"`
		Encryption TCPClientEncryptoin `mapstructure:"Encryption"`
	}
)
