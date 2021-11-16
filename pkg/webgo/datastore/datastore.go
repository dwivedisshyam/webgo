package datastore

import (
	"github.com/dwivedisshyam/webgo/pkg/log"
)

type Datastore struct {
	SQLClient *SQL
	Logger    *log.Logger
}

func (d *Datastore) SetDB(s *SQL) {
	d.SQLClient = s
}

func (d *Datastore) DB() *SQL {
	return d.SQLClient
}
