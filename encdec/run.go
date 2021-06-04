package encdec

import (
	"os"

	"github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05.999"})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logrus.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	logrus.SetLevel(logrus.WarnLevel)
}

var BaseLogger = logrus.WithFields(logrus.Fields{
	"app":    "inputYourAppName",
	"author": "inputYourName",
})

func Run() {
	logger := BaseLogger.WithFields(logrus.Fields{"component": "main"})
	msg := "hello, world"
	textPasswd := ""
	encrypted, _ := EncryptToBase64(msg, textPasswd)
	logger.Warn(encrypted)

	encrypted001 := "G0xxNVOvUCwGl811S3pMiTkl71zj_I2mgakiO0hP54USBgXPaOF98Q=="
	r, _ := DecryptFromBase64(encrypted001, textPasswd)
	logger.Warn(r)
}
