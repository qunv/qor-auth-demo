package resources

import (
	qadmin "github.com/qor/admin"
)

type Resources struct {
	Admin *qadmin.Admin
}

func AddResources(qorAdmin *qadmin.Admin) {
	r := Resources{Admin: qorAdmin}
	r.AddUser()
	r.AddProduct()
}
