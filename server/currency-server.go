package server

import (
	"context"

	protos "github.com/FadyGamilM/gogrpc/protos/currency"
	hclog "github.com/hashicorp/go-hclog"
)
type Currency struct {
	logger hclog.Logger
	protos.UnimplementedCurrencyServer 
}

func (c *Currency) GetRate(ctx context.Context,req *protos.RateReq)(*protos.RateRes, error){
	c.logger.Info("Handle GetRate", "Base : ", req.GetBaseCurrencyType(), "Destination : ", req.GetDestinationCurrencyType())

	return &protos.RateRes{Rate: float32(0.5)}, nil
}

func NewCurrencyServer (l hclog.Logger) *Currency{
	return &Currency{logger: l}
}