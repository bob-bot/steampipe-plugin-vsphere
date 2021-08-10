package vsphere

import (
	"context"
	"fmt"
	"net/url"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/vmware/govmomi/session/cache"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/soap"
)

func connect(ctx context.Context, d *plugin.QueryData) (*vim25.Client, error) {
	logger := plugin.Logger(ctx)
	vsphereConfig := GetConfig(d.Connection)
	client := new(vim25.Client)

	parsedUrl, err := soap.ParseURL(*vsphereConfig.BaseUrl)
	if err != nil {
		logger.Error(fmt.Sprintf("23: %v", err))
	}
	parsedUrl.User = url.UserPassword(*vsphereConfig.Username, *vsphereConfig.Password)

	session := &cache.Session{
		URL:      parsedUrl,
		Insecure: *vsphereConfig.Insecure,
	}
	err = session.Login(ctx, client, nil)
	if err != nil {
		logger.Error(fmt.Sprintf("34: %v", err))
	}

	return client, nil
}