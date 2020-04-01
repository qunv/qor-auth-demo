package roles

import (
	"fmt"
	"net/http"

	"github.com/qor/roles"
	"github.com/quannv132/qor/models"
)

// Role names
const (
	Admin   = "admin"
	Manager = "manager"
)

// Definition of "admin" and "not admin" roles
var (
	RolesList     = []string{Admin, Manager}
	NotAdminRoles = []string{Manager}
)

// Register roles on startup
func Load() error {

	roles.Register(Admin, func(req *http.Request, currentUser interface{}) bool {
		usr, ok := currentUser.(*models.User)
		fmt.Println("User role: ", usr.Role)
		if !ok {
			return false
		}
		return usr.Role == Admin
	})

	roles.Register(Manager, func(req *http.Request, currentUser interface{}) bool {
		usr, ok := currentUser.(*models.User)
		fmt.Println("User role: ", usr.Role)
		if !ok {
			return false
		}
		return usr.Role == Manager
	})

	return nil
}
