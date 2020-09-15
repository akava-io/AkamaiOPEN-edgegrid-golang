package appsec

import (
	"fmt"

	"github.com/akamai/AkamaiOPEN-edgegrid-golang/client-v1"
	edge "github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"
)

// RatePolicy represents a collection of RatePolicy
//
// See: RatePolicy.GetRatePolicy()
// API Docs: // appsec v1
//
// https://developer.akamai.com/api/cloud_security/application_security/v1.html

type RatePolicyResponse struct {
	ID                    int    `json:"id"`
	PolicyID              int    `json:"policyId"`
	ConfigID              int    `json:"configId"`
	ConfigVersion         int    `json:"configVersion"`
	MatchType             string `json:"matchType"`
	Type                  string `json:"type"`
	Name                  string `json:"name"`
	Description           string `json:"description"`
	AverageThreshold      int    `json:"averageThreshold"`
	BurstThreshold        int    `json:"burstThreshold"`
	ClientIdentifier      string `json:"clientIdentifier"`
	UseXForwardForHeaders bool   `json:"useXForwardForHeaders"`
	RequestType           string `json:"requestType"`
	SameActionOnIpv6      bool   `json:"sameActionOnIpv6"`
	Path                  struct {
		PositiveMatch bool     `json:"positiveMatch"`
		Values        []string `json:"values"`
	} `json:"path"`
	PathMatchType        string `json:"pathMatchType"`
	PathURIPositiveMatch bool   `json:"pathUriPositiveMatch"`
	FileExtensions       struct {
		PositiveMatch bool     `json:"positiveMatch"`
		Values        []string `json:"values"`
	} `json:"fileExtensions"`
	Hostnames              []string `json:"hostNames"`
	AdditionalMatchOptions []struct {
		PositiveMatch bool     `json:"positiveMatch"`
		Type          string   `json:"type"`
		Values        []string `json:"values"`
	} `json:"additionalMatchOptions"`
	QueryParameters []struct {
		Name          string   `json:"name"`
		Values        []string `json:"values"`
		PositiveMatch bool     `json:"positiveMatch"`
		ValueInRange  bool     `json:"valueInRange"`
	} `json:"queryParameters"`
	CreateDate string `json:"createDate"`
	UpdateDate string `json:"updateDate"`
	Used       bool   `json:"used"`
}

// NewCpCodes creates a new *CpCodes
func NewRatePolicyResponse() *RatePolicyResponse {
	return &RatePolicyResponse{}
}

// GetRatePolicy populates a *RatePolicy with it's related RatePolicy
//
// API Docs: // appsec v1
//
// https://developer.akamai.com/api/cloud_security/application_security/v1.html#getratepolicy

func (ratepolicy *RatePolicyResponse) GetRatePolicy(correlationid string) error {

	req, err := client.NewRequest(
		Config,
		"GET",
		fmt.Sprintf(
			"/appsec/v1/configs/%d/versions/%d/rate-policies/%d",
			ratepolicy.ConfigID,
			ratepolicy.ConfigVersion,
			ratepolicy.ID,
		),
		nil,
	)
	if err != nil {
		return err
	}

	edge.PrintHttpRequestCorrelation(req, true, correlationid)

	res, err := client.Do(Config, req)
	if err != nil {
		return err
	}

	edge.PrintHttpResponseCorrelation(res, true, correlationid)

	if client.IsError(res) {
		return client.NewAPIError(res)
	}

	if err = client.BodyJSON(res, ratepolicy); err != nil {
		return err
	}

	return nil

}

// Update will update a RatePolicy.
//
// API Docs: // appsec v1
//
// https://developer.akamai.com/api/cloud_security/application_security/v1.html#putratepolicy
func (ratepolicy *RatePolicyResponse) UpdateRatePolicy(correlationid string) error {
	req, err := client.NewJSONRequest(
		Config,
		"PUT",
		fmt.Sprintf(
			"/appsec/v1/configs/%d/versions/%d/rate-policies/%d",
			ratepolicy.ConfigID,
			ratepolicy.ConfigVersion,
			ratepolicy.ID,
		),
		ratepolicy,
	)
	if err != nil {
		return err
	}

	edge.PrintHttpRequestCorrelation(req, true, correlationid)

	res, err := client.Do(Config, req)
	if err != nil {
		return err
	}

	edge.PrintHttpResponseCorrelation(res, true, correlationid)

	if client.IsError(res) {
		return client.NewAPIError(res)
	}

	return nil
}

// Save will create a new ratepolicy.
//
//
// API Docs: // appsec v1
//
// https://developer.akamai.com/api/cloud_security/application_security/v1.html#postratepolicy
func (ratepolicy *RatePolicyResponse) SaveRatePolicy(correlationid string) error {
	req, err := client.NewJSONRequest(
		Config,
		"POST",
		fmt.Sprintf(
			"/appsec/v1/configs/%d/versions/%d/rate-policies",
			ratepolicy.ConfigID,
			ratepolicy.ConfigVersion,
		),
		ratepolicy,
	)
	if err != nil {
		return err
	}

	edge.PrintHttpRequestCorrelation(req, true, correlationid)

	res, err := client.Do(Config, req)
	if err != nil {
		return err
	}

	edge.PrintHttpResponseCorrelation(res, true, correlationid)

	if client.IsError(res) {
		return client.NewAPIError(res)
	}

	if err = client.BodyJSON(res, ratepolicy); err != nil {
		return err
	}

	return nil
}

// Delete will delete a RatePolicy
//
//
// API Docs: // appsec v1
//
// https://developer.akamai.com/api/cloud_security/application_security/v1.html#deleteratepolicy
func (ratepolicy *RatePolicyResponse) DeleteRatePolicy(correlationid string) error {
	req, err := client.NewJSONRequest(
		Config,
		"DELETE",
		fmt.Sprintf(
			"/appsec/v1/configs/%d/versions/%d/rate-policies/%d",
			ratepolicy.ConfigID,
			ratepolicy.ConfigVersion,
			ratepolicy.ID,
		),
		nil,
	)
	if err != nil {
		return err
	}

	edge.PrintHttpRequestCorrelation(req, true, correlationid)

	res, err := client.Do(Config, req)
	if err != nil {
		return err
	}

	edge.PrintHttpResponseCorrelation(res, true, correlationid)

	if client.IsError(res) {
		return client.NewAPIError(res)
	}

	return nil
}
