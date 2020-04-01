package migrate

import (
	"github.com/jinzhu/gorm"
	"github.com/quannv132/qor/models"
	"gopkg.in/gormigrate.v1"
)

// Start starts the migration process
func Start(db *gorm.DB) error {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		uuidCheck,
		initial,
		models.AdminUserMigration,
		productTags,
	})
	return m.Migrate()
}
