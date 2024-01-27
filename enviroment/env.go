package enviroment

type Env struct {
	hostPort string
}

func (e *Env) GetPort() string {
	return e.hostPort
}
