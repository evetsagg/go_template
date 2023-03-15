package business

type LoggingI interface {
	//todo add more functionality
	Info(string)
	Error(error)
	Debug(string)
	Fatal(error)
}
type (
	BusinessFacade struct {
		logger LoggingI
	}
)

func (b *BusinessFacade) ProcessSomething() {
	b.logger.Info("In BusinessFacade.ProcessSomething()")
}

func New(logger LoggingI) *BusinessFacade {
	return &BusinessFacade{logger: logger}

}
