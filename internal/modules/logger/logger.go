package logger

import (
	"go.uber.org/zap"
)

var logg *zap.SugaredLogger

func Init() {
	log, _ := zap.NewDevelopment()

	logg = log.Sugar()

}
func LogError(tmpl string, err error) {
	Init()
	logg.Errorf(tmpl, err)
}
