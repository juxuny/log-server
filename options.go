package log_server

type Options struct {
	CertFile string
}

func DefaultOptions() Options {
	return Options{
		CertFile: "cert/ca-cert.pem",
	}
}
