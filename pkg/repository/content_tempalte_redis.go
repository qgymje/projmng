package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"projmng/pkg/entity"
	"time"

	"github.com/go-redis/redis/v8"
)

type ContentTemplateRedis struct {
	rdb *redis.Client
}

const contentTemplatePrefix = "content_tpl:"
const contentTemplateCacheDuration = 1 * time.Hour

func (ct *ContentTemplateRedis) Create(t *entity.ContentTemplate) error {
	bytes, err := json.Marshal(ct)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := ct.rdb.SetEX(ctx, fmt.Sprintf(contentTemplatePrefix+"%d", t.Id), bytes, contentTemplateCacheDuration).Err(); err != nil {
		return err
	}

	return nil
}
