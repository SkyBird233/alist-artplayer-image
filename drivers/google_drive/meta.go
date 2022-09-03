package google_drive

import (
	"github.com/alist-org/alist/v3/internal/driver"
	"github.com/alist-org/alist/v3/internal/op"
)

type Addition struct {
	driver.RootFolderID
	RefreshToken   string `json:"refresh_token" required:"true"`
	OrderBy        string `json:"order_by" type:"string" help:"such as: folder,name,modifiedTime"`
	OrderDirection string `json:"order_direction" type:"select" options:"asc,desc"`
	ClientID       string `json:"client_id" required:"true" default:"202264815644.apps.googleusercontent.com"`
	ClientSecret   string `json:"client_secret" required:"true" default:"X4Z3ca8xfWDb1Voo-F9a7ZxJ"`
}

var config = driver.Config{
	Name:        "GoogleDrive",
	OnlyProxy:   true,
	DefaultRoot: "root",
}

func New() driver.Driver {
	return &GoogleDrive{}
}

func init() {
	op.RegisterDriver(config, New)
}