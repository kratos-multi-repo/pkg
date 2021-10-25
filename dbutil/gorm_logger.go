package dbutil

import (
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm/logger"
)

type dbWriter struct {
	log *log.Helper
}

func (l *dbWriter) Printf(format string, vs ...interface{}) {
	if len(vs) > 1 {
		if _, ok := vs[1].(*mysql.MySQLError); ok {
			l.log.Errorf(format, vs...)
			return
		}
	}

	l.log.Debugf(format, vs...)
}

func newWriter(logger log.Logger) *dbWriter {
	return &dbWriter{log: log.NewHelper(logger)}
}

func newLogger(l log.Logger) logger.Interface {
	return logger.New(
		newWriter(l),
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)
}
