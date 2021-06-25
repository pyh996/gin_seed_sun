package global

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-redis/redis"
	"github.com/minio/minio-go"
	"go.uber.org/zap"
	"go_gin/config"
	"gorm.io/gorm"
)

var (
	Lg *zap.Logger

	Settings config.ServerConfig

	Trans ut.Translator

	DB    *gorm.DB

	Redis *redis.Client

	MinioClient *minio.Client
)
