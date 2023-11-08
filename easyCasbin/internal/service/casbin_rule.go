package service

import (
	"context"
	v1 "easyCasbin/api/casbin_rule/v1"
	"easyCasbin/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CasbinRuleService struct {
	v1.UnimplementedCasbinRuleServer

	uc  *biz.CasbinRuleUseCase
	log *log.Helper
}

func NewCasbinRuleService(uc *biz.CasbinRuleUseCase, logger log.Logger) *CasbinRuleService {
	return &CasbinRuleService{uc: uc, log: log.NewHelper(logger)}
}

func (s *CasbinRuleService) GetAllSubjects(context.Context, *emptypb.Empty) (*v1.GetAllSubjectsRpl, error) {
	panic(1)
}
