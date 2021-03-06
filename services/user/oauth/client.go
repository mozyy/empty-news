package oauth

import (
	"context"
	"strconv"

	"github.com/go-oauth2/oauth2/v4"
	oauthv1 "github.com/mozyy/empty-news/proto/user/oauth/v1"
	"github.com/mozyy/empty-news/utils/db"
	"gorm.io/gorm"
)

type Client struct {
	*gorm.DB
}
type Oauth2Client struct {
	*oauthv1.ClientStore
}

func NewStoreClient() oauth2.ClientStore {
	dbGorm := db.NewGorm("e_user")

	dbGorm.AutoMigrate(&oauthv1.ClientStoreGORM{})

	// dbGorm.Create(&Oauth2Client{OAuthClientORM: pbmodel.OAuthClientORM{Secret: uuid.NewString()}})
	return &Client{dbGorm}
}

// according to the ID for the client information
func (c *Client) GetByID(ctx context.Context, id string) (oauth2.ClientInfo, error) {
	client := &oauthv1.ClientStoreGORM{}
	res := c.First(client, id)
	return &Oauth2Client{ClientStore: client.ToPB(ctx)}, res.Error
}

func (c *Oauth2Client) GetID() string {
	return strconv.FormatUint(uint64(c.ID), 10)
}
