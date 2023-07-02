package appctx

import (
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type appContext struct {
	db        *gorm.DB
	secretKey string
	redis     *redis.Client
}
type AppContext interface {
	GetMainConnection() *gorm.DB
	GetSecretKey() string
	GetCache() *redis.Client
}

func NewAppContext(db *gorm.DB, secretKey string, redis *redis.Client) *appContext {
	return &appContext{
		db:        db,
		secretKey: secretKey,
		redis:     redis,
	}
}
func (appCtx *appContext) GetMainConnection() *gorm.DB { return appCtx.db }
func (appCtx *appContext) GetSecretKey() string        { return appCtx.secretKey }
func (appCtx *appContext) GetCache() *redis.Client     { return appCtx.redis }
