package cache

import (
	"context"

	"github.com/rl404/fairy/cache"
	"github.com/rl404/hayasui/internal/domain/reaction/entity"
	"github.com/rl404/hayasui/internal/errors"
	"github.com/rl404/hayasui/internal/utils"
)

// Cache contains functions for reaction cache.
type Cache struct {
	cacher cache.Cacher
}

// New to create new reaction cache.
func New(cacher cache.Cacher) *Cache {
	return &Cache{
		cacher: cacher,
	}
}

// SetCommand to set command reaction.
func (c *Cache) SetCommand(ctx context.Context, msgID string, data entity.Command) error {
	key := utils.GetKey("msg", msgID)
	if err := c.cacher.Set(ctx, key, data); err != nil {
		return errors.Wrap(ctx, err)
	}
	return nil
}

// GetCommand to get command reaction.
func (c *Cache) GetCommand(ctx context.Context, msgID string) (data *entity.Command, _ error) {
	key := utils.GetKey("msg", msgID)
	if err := c.cacher.Get(ctx, key, &data); err != nil {
		return nil, errors.Wrap(ctx, err)
	}
	return data, nil
}
