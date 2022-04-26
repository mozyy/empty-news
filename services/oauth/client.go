package oauth

import (
	"context"
	"strconv"

	"github.com/go-oauth2/oauth2/v4"
	"github.com/mozyy/empty-news/proto/pbuser"
	"github.com/mozyy/empty-news/utils/db"
	"gorm.io/gorm"
)

type Client struct {
	*gorm.DB
}
type Oauth2Client struct {
	*pbuser.OAuthClient
}

func NewStoreClient() *Client {
	dbGorm := db.NewGorm("e_user")

	dbGorm.AutoMigrate(&pbuser.OAuthClientGORM{})

	// dbGorm.Create(&Oauth2Client{OAuthClientORM: pbmodel.OAuthClientORM{Secret: uuid.NewString()}})
	return &Client{dbGorm}
}

// according to the ID for the client information
func (c *Client) GetByID(ctx context.Context, id string) (oauth2.ClientInfo, error) {
	client := &pbuser.OAuthClientGORM{}
	res := c.First(client, id)
	return &Oauth2Client{OAuthClient: client.ToPB(ctx)}, res.Error
}

func (c *Oauth2Client) GetID() string {
	return strconv.FormatUint(uint64(c.ID), 10)
}
