package graceful

type process struct {
	starter  starter
	disabled bool
}

func NewProcess(starter starter) process {
	return process{
		starter:  starter,
		disabled: false,
	}
}

func (p process) Disable(d bool) process {
	p.disabled = d

	return p
}
