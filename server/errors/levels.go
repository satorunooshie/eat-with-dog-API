package errors

type level string

const (
	levelPanic     level = "panic"
	levelEmergency level = "emergency"
	levelAlert     level = "alert"
	levelCritical  level = "critical"
	levelError     level = "error"
	levelWarn      level = "warn"
	levelNotice    level = "notice"
	levelInfo      level = "info"
	levelDebug     level = "debug"
)

func (e *appError) Panic() AppError {
	e.level = levelPanic
	return e
}

func (e *appError) Critical() AppError {
	e.level = levelCritical
	return e
}

func (e *appError) Warn() AppError {
	e.level = levelWarn
	return e
}

func (e *appError) Info() AppError {
	e.level = levelInfo
	return e
}

func (e *appError) IsPanic() bool {
	return e.checkLevel(levelPanic)
}

func (e *appError) IsCritical() bool {
	return e.checkLevel(levelCritical)
}

func (e *appError) IsWarn() bool {
	return e.checkLevel(levelWarn)
}

func (e *appError) IsInfo() bool {
	return e.checkLevel(levelInfo)
}

func (e *appError) checkLevel(l level) bool {
	if e.level != `` {
		return e.level == l
	}
	next := AsAppError(e.next)
	if next != nil {
		next.checkLevel(l)
	}
	// default critical
	return l == levelCritical
}
