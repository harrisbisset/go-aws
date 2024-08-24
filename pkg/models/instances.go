package models

type (
	ConnectionReport struct {
		connected bool
		git       string
	}
)

func (c ConnectionReport) Connected() bool {
	return c.connected
}

func (c ConnectionReport) Git() string {
	return c.git
}
