package main

import (
	"context"

	"zgo.at/goatcounter/v2"
	"zgo.at/zdb"
)

func AutoCreateSite(ctx context.Context, db zdb.DB, cfg *Config) error {
	ctx = zdb.WithDB(ctx, db)
	// Create site
	s := goatcounter.Site{Cname: &cfg.IPAddress}
	s.Defaults(ctx)
	err := s.Insert(ctx)
	if err != nil {
		return err
	}

	// Create user
	u := goatcounter.User{
		Site:     s.ID,
		Email:    cfg.Email,
		Password: []byte(cfg.Password),
		Access:   goatcounter.UserAccesses{"all": goatcounter.AccessAdmin},
	}
	err = u.Insert(ctx, false)
	if err != nil {
		return err
	}

	return nil
}
