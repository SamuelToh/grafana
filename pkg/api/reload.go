package api

import (
	"github.com/grafana/grafana/pkg/bus"
	m "github.com/grafana/grafana/pkg/models"
)

func ReloadLdapCfg() Response {
	if err := bus.Dispatch(&m.ReloadLdapCfgCmd{}); err != nil {
		return Error(500, "Failed to reload ldap config.", err)
	}

	return Success("Ldap config reloaded")
}

func ReloadPlugins() Response {
	if err := bus.Dispatch(&m.ReloadPluginsCmd{}); err != nil {
		return Error(500, "Failed to reload plugins.", err)
	}

	return Success("Plugins reloaded")
}
