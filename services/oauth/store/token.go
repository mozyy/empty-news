package store

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/mozyy/empty-news/utils/db"
	"gorm.io/gorm"
)

// Oauth2Token data item
type Oauth2Token struct {
	gorm.Model

	ExpiredAt int64
	Code      string `gorm:"type:varchar(512)"`
	Access    string `gorm:"type:varchar(512)"`
	Refresh   string `gorm:"type:varchar(512)"`
	Data      string `gorm:"type:text"`
}

// NewStoreToken create mysql store instance,
func NewStoreToken() *StoreToken {

	dbGorm := db.NewGorm("e_user")

	return NewStoreWithDB(dbGorm)
}

func NewStoreWithDB(db *gorm.DB) *StoreToken {
	store := &StoreToken{
		db:     db,
		stdout: os.Stderr,
	}
	interval := 600

	store.ticker = time.NewTicker(time.Second * time.Duration(interval))

	if !db.Migrator().HasTable(Oauth2Token{}) {
		if err := db.Migrator().CreateTable(&Oauth2Token{}); err != nil {
			panic(err)
		}
	}

	go store.gc()
	return store
}

// StoreToken mysql token store
type StoreToken struct {
	db     *gorm.DB
	stdout io.Writer
	ticker *time.Ticker
}

// SetStdout set error output
func (s *StoreToken) SetStdout(stdout io.Writer) *StoreToken {
	s.stdout = stdout
	return s
}

// Close close the store
func (s *StoreToken) Close() {
	s.ticker.Stop()
}

func (s *StoreToken) errorf(format string, args ...interface{}) {
	if s.stdout != nil {
		buf := fmt.Sprintf(format, args...)
		s.stdout.Write([]byte(buf))
	}
}

func (s *StoreToken) gc() {
	for range s.ticker.C {
		now := time.Now().Unix()
		var count int64
		if err := s.db.Model(&Oauth2Token{}).Where("expired_at <= ?", now).Or("code = ? and access = ? AND refresh = ?", "", "", "").Count(&count).Error; err != nil {
			s.errorf("[ERROR]:%s\n", err)
			return
		}
		if count > 0 {
			// not soft delete.
			if err := s.db.Model(&Oauth2Token{}).Where("expired_at <= ?", now).Or("code = ? and access = ? AND refresh = ?", "", "", "").Unscoped().Delete(&Oauth2Token{}).Error; err != nil {
				s.errorf("[ERROR]:%s\n", err)
			}
		}
	}
}

// Create create and store the new token information
func (s *StoreToken) Create(ctx context.Context, info oauth2.TokenInfo) error {
	jv, err := json.Marshal(info)
	if err != nil {
		return err
	}
	item := &Oauth2Token{
		Data: string(jv),
	}

	if code := info.GetCode(); code != "" {
		item.Code = code
		item.ExpiredAt = info.GetCodeCreateAt().Add(info.GetCodeExpiresIn()).Unix()
	} else {
		item.Access = info.GetAccess()
		item.ExpiredAt = info.GetAccessCreateAt().Add(info.GetAccessExpiresIn()).Unix()

		if refresh := info.GetRefresh(); refresh != "" {
			item.Refresh = info.GetRefresh()
			item.ExpiredAt = info.GetRefreshCreateAt().Add(info.GetRefreshExpiresIn()).Unix()
		}
	}

	return s.db.WithContext(ctx).Create(item).Error
}

// RemoveByCode delete the authorization code
func (s *StoreToken) RemoveByCode(ctx context.Context, code string) error {
	return s.db.WithContext(ctx).
		Model(&Oauth2Token{}).
		Where("code = ?", code).
		Update("code", "").
		Error
}

// RemoveByAccess use the access token to delete the token information
func (s *StoreToken) RemoveByAccess(ctx context.Context, access string) error {
	return s.db.WithContext(ctx).
		Model(&Oauth2Token{}).
		Where("access = ?", access).
		Update("access", "").
		Error
}

// RemoveByRefresh use the refresh token to delete the token information
func (s *StoreToken) RemoveByRefresh(ctx context.Context, refresh string) error {
	return s.db.WithContext(ctx).
		Model(&Oauth2Token{}).
		Where("refresh = ?", refresh).
		Update("refresh", "").
		Error
}

func (s *StoreToken) toTokenInfo(data string) oauth2.TokenInfo {
	var tm models.Token
	err := json.Unmarshal([]byte(data), &tm)
	if err != nil {
		return nil
	}
	return &tm
}

// GetByCode use the authorization code for token information data
func (s *StoreToken) GetByCode(ctx context.Context, code string) (oauth2.TokenInfo, error) {
	if code == "" {
		return nil, nil
	}

	var item Oauth2Token
	if err := s.db.WithContext(ctx).
		Where("code = ?", code).
		Find(&item).Error; err != nil {
		return nil, err
	}
	if item.ID == 0 {
		return nil, nil
	}

	return s.toTokenInfo(item.Data), nil
}

// GetByAccess use the access token for token information data
func (s *StoreToken) GetByAccess(ctx context.Context, access string) (oauth2.TokenInfo, error) {
	if access == "" {
		return nil, nil
	}

	var item Oauth2Token
	if err := s.db.WithContext(ctx).
		Where("access = ?", access).
		Find(&item).Error; err != nil {
		return nil, err
	}
	if item.ID == 0 {
		return nil, nil
	}

	return s.toTokenInfo(item.Data), nil
}

// GetByRefresh use the refresh token for token information data
func (s *StoreToken) GetByRefresh(ctx context.Context, refresh string) (oauth2.TokenInfo, error) {
	if refresh == "" {
		return nil, nil
	}

	var item Oauth2Token
	if err := s.db.WithContext(ctx).
		Where("refresh = ?", refresh).
		Find(&item).Error; err != nil {
		return nil, err
	}
	if item.ID == 0 {
		return nil, nil
	}

	return s.toTokenInfo(item.Data), nil
}
