// Code generated by "make api"; DO NOT EDIT.
package hosts

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/kr/pretty"

	"github.com/hashicorp/boundary/api"
	"github.com/hashicorp/boundary/api/scopes"
)

type HostSet struct {
	Id          string            `json:"id,omitempty"`
	Scope       *scopes.ScopeInfo `json:"scope,omitempty"`
	Name        string            `json:"name,omitempty"`
	Description string            `json:"description,omitempty"`
	CreatedTime time.Time         `json:"created_time,omitempty"`
	UpdatedTime time.Time         `json:"updated_time,omitempty"`
	Version     uint32            `json:"version,omitempty"`
	Type        string            `json:"type,omitempty"`
	Size        uint32            `json:"size,omitempty"`
	Hosts       []*Host           `json:"hosts,omitempty"`
}

type HostSetsClient struct {
	client *api.Client
}

func NewHostSetsClient(c *api.Client) *HostSetsClient {
	return &HostSetsClient{client: c}
}

func (c *HostSetsClient) Create(ctx context.Context, hostCatalogId string, opt ...Option) (r *HostSet, apiErr error, reqErr error) {
	if hostCatalogId == "" {
		return nil, nil, fmt.Errorf("empty hostCatalogId value passed into Create request")
	}
	opts, apiOpts := getOpts(opt...)
	if c.client == nil {
		return nil, nil, fmt.Errorf("nil client")
	}

	req, err := c.client.NewRequest(ctx, "POST", fmt.Sprintf("host-catalogs/%s/host-sets", hostCatalogId), opts.valueMap, apiOpts...)
	if err != nil {
		return nil, nil, fmt.Errorf("error creating Create request: %w", err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("error performing client request during Create call: %w", err)
	}

	target := new(HostSet)
	apiErr, err = resp.Decode(target)
	if err != nil {
		return nil, nil, fmt.Errorf("error decoding Create response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr, nil
	}
	return target, apiErr, nil
}

func (c *HostSetsClient) Read(ctx context.Context, hostCatalogId string, hostSetId string, opt ...Option) (r *HostSet, apiErr error, reqErr error) {
	if hostCatalogId == "" {
		return nil, nil, fmt.Errorf("empty hostCatalogId value passed into Read request")
	}

	if hostSetId == "" {
		return nil, nil, fmt.Errorf("empty hostSetId value passed into Read request")
	}

	if c.client == nil {
		return nil, nil, fmt.Errorf("nil client")
	}

	_, apiOpts := getOpts(opt...)

	req, err := c.client.NewRequest(ctx, "GET", fmt.Sprintf("host-catalogs/%s/host-sets/%s", hostCatalogId, hostSetId), nil, apiOpts...)
	if err != nil {
		return nil, nil, fmt.Errorf("error creating Read request: %w", err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("error performing client request during Read call: %w", err)
	}

	target := new(HostSet)
	apiErr, err = resp.Decode(target)
	if err != nil {
		return nil, nil, fmt.Errorf("error decoding Read response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr, nil
	}
	return target, apiErr, nil
}

func (c *HostSetsClient) Update(ctx context.Context, hostCatalogId string, hostSetId string, version uint32, opt ...Option) (r *HostSet, apiErr error, reqErr error) {
	if hostCatalogId == "" {
		return nil, nil, fmt.Errorf("empty hostCatalogId value passed into Update request")
	}
	if hostSetId == "" {
		return nil, nil, fmt.Errorf("empty hostSetId value passed into Update request")
	}
	if c.client == nil {
		return nil, nil, fmt.Errorf("nil client")
	}

	opts, apiOpts := getOpts(opt...)

	if version == 0 {
		if !opts.withAutomaticVersioning {
			return nil, nil, errors.New("zero version number passed into Update request and automatic versioning not specified")
		}
		existingTarget, existingApiErr, existingErr := c.Read(ctx, hostCatalogId, hostSetId, opt...)
		if existingErr != nil {
			return nil, nil, fmt.Errorf("error performing initial check-and-set read: %w", existingErr)
		}
		if existingApiErr != nil {
			return nil, nil, fmt.Errorf("error from controller when performing initial check-and-set read: %s", pretty.Sprint(existingApiErr))
		}
		if existingTarget == nil {
			return nil, nil, errors.New("nil resource found when performing initial check-and-set read")
		}
		version = existingTarget.Version
	}

	opts.valueMap["version"] = version

	req, err := c.client.NewRequest(ctx, "PATCH", fmt.Sprintf("host-catalogs/%s/host-sets/%s", hostCatalogId, hostSetId), opts.valueMap, apiOpts...)
	if err != nil {
		return nil, nil, fmt.Errorf("error creating Update request: %w", err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("error performing client request during Update call: %w", err)
	}

	target := new(HostSet)
	apiErr, err = resp.Decode(target)
	if err != nil {
		return nil, nil, fmt.Errorf("error decoding Update response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr, nil
	}
	return target, apiErr, nil
}

func (c *HostSetsClient) Delete(ctx context.Context, hostCatalogId string, hostSetId string, opt ...Option) (b bool, apiErr error, reqErr error) {
	if hostCatalogId == "" {
		return false, nil, fmt.Errorf("empty hostCatalogId value passed into Delete request")
	}

	if hostSetId == "" {
		return false, nil, fmt.Errorf("empty hostSetId value passed into Delete request")
	}

	if c.client == nil {
		return false, nil, fmt.Errorf("nil client")
	}

	_, apiOpts := getOpts(opt...)

	req, err := c.client.NewRequest(ctx, "DELETE", fmt.Sprintf("host-catalogs/%s/host-sets/%s", hostCatalogId, hostSetId), nil, apiOpts...)
	if err != nil {
		return false, nil, fmt.Errorf("error creating Delete request: %w", err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return false, nil, fmt.Errorf("error performing client request during Delete call: %w", err)
	}

	type deleteResponse struct {
		Existed bool
	}
	target := &deleteResponse{}
	apiErr, err = resp.Decode(target)
	if err != nil {
		return false, nil, fmt.Errorf("error decoding Delete response: %w", err)
	}
	if apiErr != nil {
		return false, apiErr, nil
	}
	return target.Existed, apiErr, nil
}

func (c *HostSetsClient) List(ctx context.Context, hostCatalogId string, opt ...Option) (l []*HostSet, apiErr error, reqErr error) {
	if hostCatalogId == "" {
		return nil, nil, fmt.Errorf("empty hostCatalogId value passed into List request")
	}

	if c.client == nil {
		return nil, nil, fmt.Errorf("nil client")
	}

	_, apiOpts := getOpts(opt...)

	req, err := c.client.NewRequest(ctx, "GET", fmt.Sprintf("host-catalogs/%s/host-sets", hostCatalogId), nil, apiOpts...)
	if err != nil {
		return nil, nil, fmt.Errorf("error creating List request: %w", err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("error performing client request during List call: %w", err)
	}

	type listResponse struct {
		Items []*HostSet
	}
	target := &listResponse{}
	apiErr, err = resp.Decode(target)
	if err != nil {
		return nil, nil, fmt.Errorf("error decoding List response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr, nil
	}
	return target.Items, apiErr, nil
}

func (c *HostSetsClient) AddHosts(ctx context.Context, hostCatalogId string, hostSetId string, version uint32, hostIds []string, opt ...Option) (r *HostSet, apiErr error, reqErr error) {
	if hostCatalogId == "" {
		return nil, nil, fmt.Errorf("empty hostCatalogId value passed into AddHosts request")
	}
	if hostSetId == "" {
		return nil, nil, fmt.Errorf("empty hostSetId value passed into AddHosts request")
	}
	if c.client == nil {
		return nil, nil, fmt.Errorf("nil client")
	}

	opts, apiOpts := getOpts(opt...)

	if version == 0 {
		if !opts.withAutomaticVersioning {
			return nil, nil, errors.New("zero version number passed into AddHosts request")
		}
		existingTarget, existingApiErr, existingErr := c.Read(ctx, hostCatalogId, hostSetId, opt...)
		if existingErr != nil {
			return nil, nil, fmt.Errorf("error performing initial check-and-set read: %w", existingErr)
		}
		if existingApiErr != nil {
			return nil, nil, fmt.Errorf("error from controller when performing initial check-and-set read: %s", pretty.Sprint(existingApiErr))
		}
		if existingTarget == nil {
			return nil, nil, errors.New("nil resource found when performing initial check-and-set read")
		}
		version = existingTarget.Version
	}

	opts.valueMap["version"] = version

	if len(hostIds) > 0 {
		opts.valueMap["host_ids"] = hostIds
	}

	req, err := c.client.NewRequest(ctx, "POST", fmt.Sprintf("host-catalogs/%s/host-sets/%s:add-hosts", hostCatalogId, hostSetId), opts.valueMap, apiOpts...)
	if err != nil {
		return nil, nil, fmt.Errorf("error creating AddHosts request: %w", err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("error performing client request during AddHosts call: %w", err)
	}

	target := new(HostSet)
	apiErr, err = resp.Decode(target)
	if err != nil {
		return nil, nil, fmt.Errorf("error decoding AddHosts response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr, nil
	}
	return target, apiErr, nil
}

func (c *HostSetsClient) SetHosts(ctx context.Context, hostCatalogId string, hostSetId string, version uint32, hostIds []string, opt ...Option) (r *HostSet, apiErr error, reqErr error) {
	if hostCatalogId == "" {
		return nil, nil, fmt.Errorf("empty hostCatalogId value passed into SetHosts request")
	}
	if hostSetId == "" {
		return nil, nil, fmt.Errorf("empty hostSetId value passed into SetHosts request")
	}
	if c.client == nil {
		return nil, nil, fmt.Errorf("nil client")
	}

	opts, apiOpts := getOpts(opt...)

	if version == 0 {
		if !opts.withAutomaticVersioning {
			return nil, nil, errors.New("zero version number passed into SetHosts request")
		}
		existingTarget, existingApiErr, existingErr := c.Read(ctx, hostCatalogId, hostSetId, opt...)
		if existingErr != nil {
			return nil, nil, fmt.Errorf("error performing initial check-and-set read: %w", existingErr)
		}
		if existingApiErr != nil {
			return nil, nil, fmt.Errorf("error from controller when performing initial check-and-set read: %s", pretty.Sprint(existingApiErr))
		}
		if existingTarget == nil {
			return nil, nil, errors.New("nil resource found when performing initial check-and-set read")
		}
		version = existingTarget.Version
	}

	opts.valueMap["version"] = version

	if len(hostIds) > 0 {
		opts.valueMap["host_ids"] = hostIds
	} else if hostIds != nil {
		// In this function, a non-nil but empty list means clear out
		opts.valueMap["host_ids"] = nil
	}

	req, err := c.client.NewRequest(ctx, "POST", fmt.Sprintf("host-catalogs/%s/host-sets/%s:set-hosts", hostCatalogId, hostSetId), opts.valueMap, apiOpts...)
	if err != nil {
		return nil, nil, fmt.Errorf("error creating SetHosts request: %w", err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("error performing client request during SetHosts call: %w", err)
	}

	target := new(HostSet)
	apiErr, err = resp.Decode(target)
	if err != nil {
		return nil, nil, fmt.Errorf("error decoding SetHosts response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr, nil
	}
	return target, apiErr, nil
}

func (c *HostSetsClient) RemoveHosts(ctx context.Context, hostCatalogId string, hostSetId string, version uint32, hostIds []string, opt ...Option) (r *HostSet, apiErr error, reqErr error) {
	if hostCatalogId == "" {
		return nil, nil, fmt.Errorf("empty hostCatalogId value passed into RemoveHosts request")
	}
	if hostSetId == "" {
		return nil, nil, fmt.Errorf("empty hostSetId value passed into RemoveHosts request")
	}
	if c.client == nil {
		return nil, nil, fmt.Errorf("nil client")
	}

	opts, apiOpts := getOpts(opt...)

	if version == 0 {
		if !opts.withAutomaticVersioning {
			return nil, nil, errors.New("zero version number passed into RemoveHosts request")
		}
		existingTarget, existingApiErr, existingErr := c.Read(ctx, hostCatalogId, hostSetId, opt...)
		if existingErr != nil {
			return nil, nil, fmt.Errorf("error performing initial check-and-set read: %w", existingErr)
		}
		if existingApiErr != nil {
			return nil, nil, fmt.Errorf("error from controller when performing initial check-and-set read: %s", pretty.Sprint(existingApiErr))
		}
		if existingTarget == nil {
			return nil, nil, errors.New("nil resource found when performing initial check-and-set read")
		}
		version = existingTarget.Version
	}

	opts.valueMap["version"] = version

	if len(hostIds) > 0 {
		opts.valueMap["host_ids"] = hostIds
	}

	req, err := c.client.NewRequest(ctx, "POST", fmt.Sprintf("host-catalogs/%s/host-sets/%s:remove-hosts", hostCatalogId, hostSetId), opts.valueMap, apiOpts...)
	if err != nil {
		return nil, nil, fmt.Errorf("error creating RemoveHosts request: %w", err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("error performing client request during RemoveHosts call: %w", err)
	}

	target := new(HostSet)
	apiErr, err = resp.Decode(target)
	if err != nil {
		return nil, nil, fmt.Errorf("error decoding RemoveHosts response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr, nil
	}
	return target, apiErr, nil
}
