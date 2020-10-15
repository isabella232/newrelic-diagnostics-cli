package config

import (
	"github.com/newrelic/NrDiag/logger"
	"github.com/newrelic/NrDiag/tasks"
)

// RegisterWith - will register any plugins in this package
func RegisterWith(registrationFunc func(tasks.Task, bool)) {
	logger.Debug("Registering DotNetCore/Config/*")

	registrationFunc(DotNetCoreConfigAgent{}, true)
}
