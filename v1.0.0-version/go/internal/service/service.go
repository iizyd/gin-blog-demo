package service

import (
	"context"

	"github.com/iizyd/xigua-blog/global"
	"github.com/iizyd/xigua-blog/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(global.DBEngine)
	return svc
}
