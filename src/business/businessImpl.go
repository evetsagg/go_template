package business

import (
	"go_template/src/logger"
)

type(
	BusinessFacade struct {
		logger logger.LoggingI
	}
)

func (b *BusinessFacade) ProcessSomething() {
	b.logger.Info("In BusinessFacade.ProcessSomething()")
}

func New(logger logger.LoggingI) *BusinessFacade{
	return &BusinessFacade{logger: logger}

}