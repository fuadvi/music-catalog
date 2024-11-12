package configs

type (
	Config struct {
		Service  Service  `mapstructure:"service"`
		Database Database `mapstructure:"database"`
	}

	Service struct {
		Port      string `mapstcuture:"port"`
		SecretJwt string `mapstcuture:"secretJwt"`
	}

	Database struct {
		DataSourceName string `mapstructure:"dataSourceName"`
	}
)
