package zapclilogger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

var cliEncoderPool = buffer.NewPool()

type cliEncoder struct {
	zapcore.Encoder

	verbose bool
}

func (e *cliEncoder) Clone() zapcore.Encoder {
	return &cliEncoder{
		Encoder: e.Encoder.Clone(),

		verbose: e.verbose,
	}
}

func (e *cliEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	if !e.verbose && entry.Level < zap.WarnLevel {
		buf := cliEncoderPool.Get()
		if _, err := buf.WriteString(entry.Message + "\n"); err != nil {
			return buf, err
		}

		return buf, nil
	}

	return e.Encoder.EncodeEntry(entry, fields)
}

func New(verbose bool) (*zap.Logger, error) {
	zap.RegisterEncoder("cli", func(ec zapcore.EncoderConfig) (zapcore.Encoder, error) {
		return &cliEncoder{
			Encoder: zapcore.NewConsoleEncoder(ec),

			verbose: verbose,
		}, nil
	})

	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.LevelKey = ""
	encoderConfig.CallerKey = ""
	encoderConfig.TimeKey = ""

	logConfig := zap.NewDevelopmentConfig()
	logConfig.EncoderConfig = encoderConfig
	logConfig.Encoding = "cli"

	logConfig.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	if verbose {
		logConfig.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	}

	return logConfig.Build()
}
