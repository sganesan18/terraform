package cloudflare

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// ListWAFPackages returns a slice of the WAF packages for the given zone.
func (api *API) ListWAFPackages(zoneID string) ([]WAFPackage, error) {
	var p WAFPackagesResponse
	var packages []WAFPackage
	var res []byte
	var err error
	uri := "/zones/" + zoneID + "/firewall/waf/packages"
	res, err = api.makeRequest("GET", uri, nil)
	if err != nil {
		return []WAFPackage{}, errors.Wrap(err, errMakeRequestError)
	}
	err = json.Unmarshal(res, &p)
	if err != nil {
		return []WAFPackage{}, errors.Wrap(err, errUnmarshalError)
	}
	if !p.Success {
		// TODO: Provide an actual error message instead of always returning nil
		return []WAFPackage{}, err
	}
	for pi := range p.Result {
		packages = append(packages, p.Result[pi])
	}
	return packages, nil
}

// ListWAFRules returns a slice of the WAF rules for the given WAF package.
func (api *API) ListWAFRules(zoneID, packageID string) ([]WAFRule, error) {
	var r WAFRulesResponse
	var rules []WAFRule
	var res []byte
	var err error
	uri := "/zones/" + zoneID + "/firewall/waf/packages/" + packageID + "/rules"
	res, err = api.makeRequest("GET", uri, nil)
	if err != nil {
		return []WAFRule{}, errors.Wrap(err, errMakeRequestError)
	}
	err = json.Unmarshal(res, &r)
	if err != nil {
		return []WAFRule{}, errors.Wrap(err, errUnmarshalError)
	}
	if !r.Success {
		// TODO: Provide an actual error message instead of always returning nil
		return []WAFRule{}, err
	}
	for ri := range r.Result {
		rules = append(rules, r.Result[ri])
	}
	return rules, nil
}
