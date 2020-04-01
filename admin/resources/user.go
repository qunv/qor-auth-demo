package resources

import (
	"github.com/qor/admin"
	"github.com/qor/qor"
	"github.com/qor/qor/resource"
	qroles "github.com/qor/roles"
	"github.com/qor/validations"
	"github.com/quannv132/qor/admin/permission"
	"github.com/quannv132/qor/models"
	"github.com/quannv132/qor/roles"
	"golang.org/x/crypto/bcrypt"
)

func (r Resources) AddUser() {
	usr := r.Admin.AddResource(
		&models.User{},
		&admin.Config{
			Menu:       []string{"User Management"},
			Permission: permission.ADMIN,
		},
	)
	usr.IndexAttrs("-Password")
	usr.Meta(&admin.Meta{
		Name: "Password",
		Type: "password",
		Setter: func(resource interface{}, metaValue *resource.MetaValue, context *qor.Context) {
			values := metaValue.Value.([]string)
			if len(values) > 0 {
				if np := values[0]; np != "" {
					pwd, err := bcrypt.GenerateFromPassword([]byte(np), bcrypt.DefaultCost)
					if err != nil {
						context.DB.AddError(validations.NewError(usr, "Password", "Can't encrypt password")) // nolint: gosec,errcheck
						return
					}
					u := resource.(*models.User)
					u.Password = pwd
				}
			}
		},
		Permission: qroles.Allow(qroles.CRUD, roles.Admin),
	})

	usr.Meta(&admin.Meta{
		Name: "Role",
		Config: &admin.SelectOneConfig{
			Collection: roles.RolesList,
		},
	})
}
