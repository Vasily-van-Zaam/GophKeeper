package component

type Component interface {
	HandlerError(func(msg string)) Component
	SetHandler(func()) Component
}
