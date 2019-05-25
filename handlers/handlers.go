package handlers

import (
	"github.com/caarlos0/env"
	errors "github.com/go-openapi/errors"

	models "github.com/xbcsmith/spoon/models"
	"sync"
	"sync/atomic"
)

// the variables we need throughout our implementation
var items = make(map[int64]*models.Item)
var lastID int64

// itemsLock is the lock for the items
var itemsLock = &sync.Mutex{}

// newItemId generates a uniq id for new items
func newItemID() int64 {
	return atomic.AddInt64(&lastID, 1)
}

// AddItem creates a new item
func AddItem(item *models.Item) error {
	if item == nil {
		return errors.New(500, "item must be present")
	}

	itemsLock.Lock()
	defer itemsLock.Unlock()

	newID := newItemID()
	item.ID = newID
	items[newID] = item

	return nil
}

// UpdateItem updates a single item
func UpdateItem(id int64, item *models.Item) error {
	if item == nil {
		return errors.New(500, "item must be present")
	}

	itemsLock.Lock()
	defer itemsLock.Unlock()

	_, exists := items[id]
	if !exists {
		return errors.NotFound("not found: item %d", id)
	}

	item.ID = id
	items[id] = item
	return nil
}

// GetItem returns a single item
func GetItem(id int64) (item *models.Item, err error) {
	itemsLock.Lock()
	defer itemsLock.Unlock()
	item, exists := items[id]
	if exists {
		return item, nil
	}
	return nil, errors.NotFound("not found: item %d", id)
}

// DeleteItem removes a single item from items
func DeleteItem(id int64) error {
	itemsLock.Lock()
	defer itemsLock.Unlock()

	_, exists := items[id]
	if !exists {
		return errors.NotFound("not found: item %d", id)
	}

	delete(items, id)
	return nil
}

// AllItems returns all the items
func AllItems(since int64, limit int32) (results []*models.Item) {
	results = make([]*models.Item, 0)
	itemsLock.Lock()
	defer itemsLock.Unlock()
	for id, item := range items {
		if len(results) >= int(limit) {
			return
		}
		if since == 0 || id > since {
			results = append(results, item)
		}
	}
	return
}

// LivenessCheck returns true
func LivenessCheck() (check *models.Liveness, err error) {
	live := true
	return &models.Liveness{Live: &live}, nil
}

// ReadinessCheck returns true
func ReadinessCheck() (check *models.Readiness, err error) {
	ready := true
	return &models.Readiness{Ready: &ready}, nil
}

// SpoonDbCfg config for the database connections
type SpoonDbCfg struct {
	Host   string `env:"SPOON_DATABASE_HOSTNAME" envDefault:"0.0.0.0"`
	Port   string `env:"SPOON_DATABASE_PORT" envDefault:"5432"`
	User   string `env:"SPOON_DATABASE_USER" envDefault:"postgres"`
	Passwd string `env:"SPOON_DATABASE_PASSWD" envDefault:"admin"`
	Name   string `env:"SPOON_DATABASE_NAME" envDefault:"spoons"`
}

// GetCfg get the cfg from os.Environ
func GetDbCfg() (*SpoonDbCfg, error) {
	cfg := SpoonDbCfg{}
	err := env.Parse(&cfg)
	if err != nil {
		return nil, errors.New(500, "Error reading db cfg from ENV")
	}
	return &cfg, nil
}

// SpoonCfg Cfg for app
type SpoonCfg struct {
	Url   string `env:"SPOON_URL" envDefault:"0.0.0.0"`
	Token string `env:"SPOON_TOKEN" envDefault:"1234567890"`
	Port  string `env:"SPOON_PORT" envDefault:"30570"`
	State string `env:"SPOON_STATE" envDefault:"stateless"`
}

// GetCfg get the cfg from os.Environ
func GetCfg() (*SpoonCfg, error) {
	cfg := SpoonCfg{}
	err := env.Parse(&cfg)
	if err != nil {
		return nil, errors.New(500, "Error reading cfg from ENV")
	}
	return &cfg, nil
}

func TokenAuth(token string) (*models.Principal, error) {
	cfg, err := GetCfg()
	if err == nil {
		if token == cfg.Token {
			prin := models.Principal(token)
			return &prin, nil
		}
	}
	return nil, errors.New(401, "incorrect api key auth")
}
