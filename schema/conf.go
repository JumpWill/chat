package schema

type ServerConfig struct {
	Name  string      `mapstructure:"name"`
	Port  int         `mapstructure:"port"`
	Log   string      `mapstructure:"log"`
	Host  string      `mapstructure:"host"`
	Mysql MysqlConfig `mapstructure:"mysql"`
	Redis RedisConfig `mapstructure:"redis"`
	Jwt   JwtConfig   `mapstructure:"jwt"`
}

type MysqlConfig struct {
	Host        string `mapstructure:"host"`
	Port        int    `mapstructure:"port"`
	User        string `mapstructure:"user"`
	Password    string `mapstructure:"password"`
	DataBase    string `mapstructure:"database"`
	TablePrefix string `mapstructure:"tableprefix"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB int `mapstructure:"db"`

}

type JwtConfig struct {
	Cookie string `mapstructure:"jwt_in_cookie"`
	Header string `mapstructure:"jwt_in_header"`
	Query  string `mapstructure:"jwt_in_query"`
	Secret string `mapstructure:"secret"`
}
