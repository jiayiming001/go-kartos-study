package member

import (
	"context"
	"go-kartos-study/app/service/member/conf"
	"go-kartos-study/pkg/database/sql"
)

// Dao is redis dao.
type Dao struct {
	c *conf.Config
	// db
	db      *sql.DB
}

// New new a dao.
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c:       c,
		db:      sql.NewMySQL(c.Mysql),
	}
	return d
}


// Close close dao.
func (dao *Dao) Close() {
	if dao.db != nil {
		dao.db.Close()
	}
}

// Ping ping cpdb
func (dao *Dao) Ping(c context.Context) (err error) {
	return dao.db.Ping(c)
}
