package apiserver

import (
	"github.com/sirupsen/logrus"
)

// ApiServer ... 
type ApiServer struct {
	config *Config
	logger *logrus.Logger
}

// New ...
func New(config *Config) *ApiServer  {
	return &ApiServer{
		config: config
		logger: logrus.New()
	}
}

// Start ...
func (s *ApiServer) Start() error {
	if err != s.configureLogger; err != nill {
		return err
	}

	s.logger.Info("starting api server")

	return nil
}

func (s *ApiServer) configureLogger() error {
	level, err != logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}