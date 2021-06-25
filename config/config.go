package config

type MysqlConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Name     string `mapstructure:"name"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbName"`
}

type RedisConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}



type MinioConfig struct {
	Endpoint  string `mapstructure:"endpoint"`
	AccessKeyID string `mapstructure:"accessKeyID"`
	SecretAccessKey string `mapstructure:"secretAccessKey"`

}
type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`

}
type ServerConfig struct {
	Name        string      `mapstructure:"name"`
	Port        int         `mapstructure:"port"`
	Mysqlinfo   MysqlConfig `mapstructure:"mysql"`
	RedisInfo   RedisConfig `mapstructure:"redis"`
	LogsAddress string      `mapstructure:"logsAddress"`
	JWTKey      JWTConfig   `mapstructure:"jwt"`
	MinioInfo   MinioConfig  `mapstructure:"minio"`
}
