package favorite

import (
	"context"
	"go-kartos-study/app/service/member/conf"
	"go-kartos-study/app/service/member/internal/model"
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
func (d *Dao) Close() {
	if d.db != nil {
		d.db.Close()
	}
}

// Ping ping cpdb
func (d *Dao) Ping(c context.Context) (err error) {
	return d.db.Ping(c)
}


func (dao *Dao) GetFavoriteByID(ctx context.Context, id int64) (m *model.Favorite, err error) {
	return dao.dbGetFavoriteByID(ctx,id)
}