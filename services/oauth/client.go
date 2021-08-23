package oauth

import (
	"context"
	"log"
	"strconv"

	"github.com/go-oauth2/oauth2/v4"
	"github.com/mozyy/empty-news/utils/db"
	"gorm.io/gorm"
)

type Client struct {
	*gorm.DB
}
type Oauth2Client struct {
	gorm.Model
	Secret string
	Domain string
	UserID string
}

func NewClient() *Client {
	dbGorm := db.NewGorm("e_user")

	store := &Client{
		dbGorm,
	}

	// dbGorm.Create(&Oauth2Client{Secret: "22222222", Domain: "http://localhost:9094", UserID: "222222"})
	if !dbGorm.Migrator().HasTable(Oauth2Client{}) {
		if err := dbGorm.Migrator().CreateTable(&Oauth2Client{}); err != nil {
			panic(err)
		}
	}
	return store
}

// according to the ID for the client information
func (c *Client) GetByID(ctx context.Context, id string) (oauth2.ClientInfo, error) {
	client := &Oauth2Client{}
	res := c.First(client, id)
	log.Println(res)
	return client, res.Error
}
func (c *Oauth2Client) GetID() string {
	return strconv.FormatUint(uint64(c.ID), 10)
}

func (c *Oauth2Client) GetSecret() string {
	return c.Secret
}

func (c *Oauth2Client) GetDomain() string {
	return c.Domain
}

func (c *Oauth2Client) GetUserID() string {
	return c.UserID
}
