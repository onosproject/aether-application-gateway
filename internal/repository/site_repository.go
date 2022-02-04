// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
//

package repository

import (
	"encoding/json"
	"fmt"
	"github.com/onosproject/aether-application-gateway/internal/config"
	"io/ioutil"
	"log"
	"net/http"
)

// SiteRepository -
type SiteRepository struct {
	HTTPClient *http.Client
	Roc        *config.RocConfig
}

// NewSiteRepository -> creates a new site repository
func NewSiteRepository(c *http.Client, roc *config.RocConfig) SiteRepository {
	return SiteRepository{
		HTTPClient: c,
		Roc:        roc,
	}
}

// GetSite - Get info for a site including all devices
func (c *SiteRepository) GetSite(enterpriseID, siteID string) (Site, error) {
	URL := fmt.Sprintf("%s/aether/v2.0.0/%s/enterprises/enterprise/%s/site/%s", c.Roc.Base, c.Roc.Target, enterpriseID, siteID)

	resp, err := c.HTTPClient.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Site{}, err
	}

	var site Site
	err = json.Unmarshal(body, &site)
	if err != nil {
		return Site{}, err
	}

	return site, nil
}
