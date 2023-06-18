package module

import (
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	fx.Provide(logrus.New),
)
