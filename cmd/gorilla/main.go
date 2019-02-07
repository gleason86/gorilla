package main

import (
	"github.com/1dustindavis/gorilla/pkg/catalog"
	"github.com/1dustindavis/gorilla/pkg/config"
	"github.com/1dustindavis/gorilla/pkg/gorillalog"
	"github.com/1dustindavis/gorilla/pkg/manifest"
	"github.com/1dustindavis/gorilla/pkg/process"
	"github.com/1dustindavis/gorilla/pkg/report"
)

func main() {

	// Get our configuration
	cfg := config.Get()

	// Create a new logger object
	gorillalog.NewLog(cfg)

	// Start creating GorillaReport
	report.Start()

	// Get the manifests
	gorillalog.Info("Retrieving manifest:", cfg.Manifest)
	manifests := manifest.Get()

	// Get the catalog
	gorillalog.Info("Retrieving catalog:", cfg.Catalogs)
	catalog := catalog.Get()

	// Process the manifests into install type groups
	gorillalog.Info("Processing manifest...")
	installs, uninstalls, updates := process.Manifests(manifests, catalog)

	// Prepare and install
	gorillalog.Info("Processing managed installs...")
	process.Installs(installs, catalog)

	// Prepare and uninstall
	gorillalog.Info("Processing managed uninstalls...")
	process.Uninstalls(uninstalls, catalog)

	// Prepare and update
	gorillalog.Info("Processing managed updates...")
	process.Updates(updates, catalog)

	// Save GorillaReport to disk
	gorillalog.Info("Saving GorillReport.json...")
	report.End()

	// Run CleanUp to delete old cached items and empty directories
	gorillalog.Info("Cleaning up the cache...")
	process.CleanUp()

	gorillalog.Info("Done!")
}
