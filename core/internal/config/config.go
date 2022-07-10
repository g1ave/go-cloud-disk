package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf

	Database struct {
		DSN   string
		Redis string
	}
	Email struct {
		Username string
		Password string
	}
	COS struct {
		SecretId  string
		SecretKey string
	}
}
