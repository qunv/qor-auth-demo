package permission

import (
	qroles "github.com/qor/roles"
	"github.com/quannv132/qor/roles"
)

var (
	ADMIN   = qroles.Allow(qroles.CRUD, roles.Admin)
	MANAGER = qroles.Allow(qroles.CRUD, roles.Manager)
)
