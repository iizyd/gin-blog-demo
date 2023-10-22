package dao

import "github.com/iizyd/xigua-blog/internal/model"

func (d *Dao) GetAuth(UserName, Password string) (model.Auth, error) {
	auth := model.Auth{UserName: UserName, Password: Password}
	return auth.Get(d.engine)
}
