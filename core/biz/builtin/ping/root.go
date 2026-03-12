package ping

type Command struct {
	Name   string
	Caller string
}

func (r *Command) GetName() string {
	return r.Name
}

func (r *Command) GetCaller() string {
	return r.Caller
}
