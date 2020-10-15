package config

import (
	log "github.com/newrelic/NrDiag/logger"
	"github.com/newrelic/NrDiag/tasks"
)

// RegisterWith - will register any plugins in this package
func RegisterWith(registrationFunc func(tasks.Task, bool)) {
	log.Debug("Registering Java/Config/*")

	registrationFunc(JavaConfigAgent{}, true)
	registrationFunc(JavaConfigValidate{}, true)
	registrationFunc(JavaConfigValidateSettings{}, true)
}
