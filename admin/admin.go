package admin

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/qor/admin"
	"github.com/sirupsen/logrus"

	"github.com/quannv132/qor/admin/bindatafs"
	"github.com/quannv132/qor/admin/resources"
	"github.com/quannv132/qor/roles"
)

// Admin abstracts the whole QOR Admin + authentication process
type Admin struct {
	db        *gorm.DB
	auth      auth
	adm       *admin.Admin
	adminpath string
	prefix    string
}

func New(db *gorm.DB, prefix, cookiesecret string) *Admin {
	if e := roles.Load(); e != nil {
		log.Fatal(e)
	}
	adminpath := filepath.Join(prefix, "/admin")
	a := Admin{
		db:        db,
		prefix:    prefix,
		adminpath: adminpath,
		auth: auth{
			db: db,
			paths: pathConfig{
				admin:  adminpath,
				login:  filepath.Join(prefix, "/login"),
				logout: filepath.Join(prefix, "/logout"),
			},
			session: sessionConfig{
				key:   "userid",
				name:  "admsession",
				store: cookie.NewStore([]byte(cookiesecret)),
			},
		},
	}
	a.adm = admin.New(&admin.AdminConfig{
		SiteName: "My Admin Interface",
		DB:       db,
		Auth:     a.auth,
		AssetFS:  bindatafs.AssetFS.NameSpace("admin"),
	})
	resources.AddResources(a.adm)
	return &a
}

// Bind will bind the admin interface to an already existing gin router
// (*gin.Engine).
func (a Admin) Bind(r *gin.Engine) {
	mux := http.NewServeMux()
	a.adm.MountTo(a.adminpath, mux)

	lfs := bindatafs.AssetFS.NameSpace("login")
	lfs.RegisterPath("admin/templates/")
	logintpl, err := lfs.Asset("login.html")
	if err != nil {
		logrus.WithError(err).Fatal("Unable to find HTML template for login page in admin")
	}
	r.SetHTMLTemplate(template.Must(template.New("login.html").Parse(string(logintpl))))

	g := r.Group(a.prefix)
	// fmt.Println("-------->session" + a.auth.session.key)
	g.Use(sessions.Sessions(a.auth.session.name, a.auth.session.store))
	{
		g.Any("/admin/*resources", gin.WrapH(mux))
		g.GET("/login", a.auth.GetLogin)
		g.POST("/login", a.auth.PostLogin)
		g.GET("/logout", a.auth.GetLogout)
	}
}
