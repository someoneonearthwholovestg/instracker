package instagram

import (
	"Instracker/internal/config"
	"github.com/ahmdrz/goinsta/v2"
)

// NewInstagram authorizes and returns an instance of Instagram
func NewInstagram(cfg *config.Config) (inst *Instagram, err error) {
	inst = new(Instagram)
	err = inst.setAccount(cfg.Instagram.Username, cfg.Instagram.Password)
	return
}

// Instagram is a structure for interaction with instagram account
type Instagram struct {
	Account *goinsta.Instagram
}

// SetAccount authorizes an instagram user
func (i *Instagram) setAccount(username, password string) (err error) {
	i.Account = goinsta.New(username, password)
	err = i.Account.Login()
	return
}

// GetUserByUsername return instagram user by username
func (i *Instagram) GetUserByUsername(username string) (user *goinsta.User, err error) {
	user, err = i.Account.Profiles.ByName(username)
	return
}
