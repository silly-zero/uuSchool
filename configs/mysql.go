package config

type Mysql struct {
	Address  string `yaml:"address"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DB       string `yaml:"db"`
	Param    string `yaml:"param"`
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Address + ":" + m.Port + ")/" + m.DB + "?" + m.Param
}
