package logx

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestNewLogger(t *testing.T) {
	l := NewLogger()

	l.Debug("Debug")
	l.Info("Info")
	l.Warn("Warn")
	l.Error("Error")

	l.Debugf("Debugf")
	l.Infof("Infof")
	l.Warnf("Warnf")
	l.Errorf("Errorf")

	l.Debugln("Debugln")
	l.Infoln("Infoln")
	l.Warnln("Warnln")
	l.Errorln("Errorln")

	l.WithContext(context.Background()).Info("WithContext")
	l.WithField("k", "v").Info("WithField")
	l.WithTime(time.Now().AddDate(-1, 0, 0)).Info("WithTime")
	l.WithError(errors.New("error")).Info("WithError")
	l.WithFields(map[string]interface{}{"k1": "v1", "k2": "v2"}).Info("WithFields")
}

func TestDefaultLogger(t *testing.T) {
	t.Log(DefaultLogger().GetLevel())

	SetLevel(WarnLevel)

	Debug("Debug")
	Info("Info")
	Warn("Warn")
	Error("Error")

	Debugf("Debugf")
	Infof("Infof")
	Warnf("Warnf")
	Errorf("Errorf")

	Debugln("Debugln")
	Infoln("Infoln")
	Warnln("Warnln")
	Errorln("Errorln")

	WithContext(context.Background()).Info("WithContext")
	WithField("k", "v").Info("WithField")
	WithTime(time.Now().AddDate(-1, 0, 0)).Info("WithTime")
	WithError(errors.New("error")).Info("WithError")
	WithFields(map[string]interface{}{"k1": "v1", "k2": "v2"}).Info("WithFields")
}

func TestDiscardLogger(t *testing.T) {
	l := DiscardLogger()

	l.Debug("Debug")
	l.Info("Info")
	l.Warn("Warn")
	l.Error("Error")

	l.Debugf("Debugf")
	l.Infof("Infof")
	l.Warnf("Warnf")
	l.Errorf("Errorf")

	l.Debugln("Debugln")
	l.Infoln("Infoln")
	l.Warnln("Warnln")
	l.Errorln("Errorln")

	l.WithContext(context.Background()).Info("WithContext")
	l.WithField("k", "v").Info("WithField")
	l.WithTime(time.Now().AddDate(-1, 0, 0)).Info("WithTime")
	l.WithError(errors.New("error")).Info("WithError")
	l.WithFields(map[string]interface{}{"k1": "v1", "k2": "v2"}).Info("WithFields")
}
