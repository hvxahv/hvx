package acct


type ServerConfig struct {
	Port string
}

func NewServerConfig(port string) *ServerConfig {
	return &ServerConfig{Port: port}
}


func (sc *ServerConfig) Start() {

}
