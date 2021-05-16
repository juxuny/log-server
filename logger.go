package log_server

type Logger interface {
	Info(app string, msg string) error
}
