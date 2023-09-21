package tracer

type Option func(*tracer)

func WithEnvironment(env string) Option {
	return func(s *tracer) {
		s.environment = env
	}
}

func WithAppName(appName string) Option {
	return func(s *tracer) {
		s.appName = appName
	}
}

func WithServiceName(serviceName string) Option {
	return func(s *tracer) {
		s.serviceName = serviceName
	}
}

func WithServerName(serverName string) Option {
	return func(s *tracer) {
		s.serverName = serverName
	}
}

func WithLanguage(language string) Option {
	return func(s *tracer) {
		s.language = language
	}
}
