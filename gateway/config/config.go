package config

type Config struct {
	Server Server `mapstructure:"server"`
	Mysql  Mysql  `mapstructure:"mysql"`
	Etcd   Etcd   `mapstructure:"etcd"`
	Redis  Redis  `mapstructure:"redis"`
	Domain Domain `mapstructure:"domain"`
}

type Server struct {
	Domain      string `mapstructure:"domain"`
	Version     string `mapstructure:"version"`
	JwtSecret   string `mapstructure:"jwtSecret"`
	GrpcAddress string `mapstructure:"grpcAddress"`
	Port        string `mapstructure:"port"`
}

type Mysql struct {
	DriverName string `mapstructure:"driverName"`
	Host       string `mapstructure:"host"`
	Port       string `mapstructure:"port"`
	Database   string `mapstructure:"database"`
	Username   string `mapstructure:"username"`
	Password   string `mapstructure:"password"`
	Charset    string `mapstructure:"charset"`
}

type Etcd struct {
	Address string `mapstructure:"address"`
}

type Redis struct {
	Address  string `mapstructure:"address"`
	Password string `mapstructure:"password"`
}

type Domain struct {
	User string `mapstructure:"user"`
	Task string `mapstructure:"task"`
}
