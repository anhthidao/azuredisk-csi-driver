//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5"
	"net/http"
	"net/url"
	"regexp"
)

// VirtualMachineScaleSetExtensionsServer is a fake server for instances of the armcompute.VirtualMachineScaleSetExtensionsClient type.
type VirtualMachineScaleSetExtensionsServer struct {
	// BeginCreateOrUpdate is the fake for method VirtualMachineScaleSetExtensionsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, vmScaleSetName string, vmssExtensionName string, extensionParameters armcompute.VirtualMachineScaleSetExtension, options *armcompute.VirtualMachineScaleSetExtensionsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armcompute.VirtualMachineScaleSetExtensionsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method VirtualMachineScaleSetExtensionsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, vmScaleSetName string, vmssExtensionName string, options *armcompute.VirtualMachineScaleSetExtensionsClientBeginDeleteOptions) (resp azfake.PollerResponder[armcompute.VirtualMachineScaleSetExtensionsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method VirtualMachineScaleSetExtensionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, vmScaleSetName string, vmssExtensionName string, options *armcompute.VirtualMachineScaleSetExtensionsClientGetOptions) (resp azfake.Responder[armcompute.VirtualMachineScaleSetExtensionsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method VirtualMachineScaleSetExtensionsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, vmScaleSetName string, options *armcompute.VirtualMachineScaleSetExtensionsClientListOptions) (resp azfake.PagerResponder[armcompute.VirtualMachineScaleSetExtensionsClientListResponse])

	// BeginUpdate is the fake for method VirtualMachineScaleSetExtensionsClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginUpdate func(ctx context.Context, resourceGroupName string, vmScaleSetName string, vmssExtensionName string, extensionParameters armcompute.VirtualMachineScaleSetExtensionUpdate, options *armcompute.VirtualMachineScaleSetExtensionsClientBeginUpdateOptions) (resp azfake.PollerResponder[armcompute.VirtualMachineScaleSetExtensionsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewVirtualMachineScaleSetExtensionsServerTransport creates a new instance of VirtualMachineScaleSetExtensionsServerTransport with the provided implementation.
// The returned VirtualMachineScaleSetExtensionsServerTransport instance is connected to an instance of armcompute.VirtualMachineScaleSetExtensionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewVirtualMachineScaleSetExtensionsServerTransport(srv *VirtualMachineScaleSetExtensionsServer) *VirtualMachineScaleSetExtensionsServerTransport {
	return &VirtualMachineScaleSetExtensionsServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armcompute.VirtualMachineScaleSetExtensionsClientCreateOrUpdateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armcompute.VirtualMachineScaleSetExtensionsClientDeleteResponse]](),
		newListPager:        newTracker[azfake.PagerResponder[armcompute.VirtualMachineScaleSetExtensionsClientListResponse]](),
		beginUpdate:         newTracker[azfake.PollerResponder[armcompute.VirtualMachineScaleSetExtensionsClientUpdateResponse]](),
	}
}

// VirtualMachineScaleSetExtensionsServerTransport connects instances of armcompute.VirtualMachineScaleSetExtensionsClient to instances of VirtualMachineScaleSetExtensionsServer.
// Don't use this type directly, use NewVirtualMachineScaleSetExtensionsServerTransport instead.
type VirtualMachineScaleSetExtensionsServerTransport struct {
	srv                 *VirtualMachineScaleSetExtensionsServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armcompute.VirtualMachineScaleSetExtensionsClientCreateOrUpdateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armcompute.VirtualMachineScaleSetExtensionsClientDeleteResponse]]
	newListPager        *tracker[azfake.PagerResponder[armcompute.VirtualMachineScaleSetExtensionsClientListResponse]]
	beginUpdate         *tracker[azfake.PollerResponder[armcompute.VirtualMachineScaleSetExtensionsClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for VirtualMachineScaleSetExtensionsServerTransport.
func (v *VirtualMachineScaleSetExtensionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "VirtualMachineScaleSetExtensionsClient.BeginCreateOrUpdate":
		resp, err = v.dispatchBeginCreateOrUpdate(req)
	case "VirtualMachineScaleSetExtensionsClient.BeginDelete":
		resp, err = v.dispatchBeginDelete(req)
	case "VirtualMachineScaleSetExtensionsClient.Get":
		resp, err = v.dispatchGet(req)
	case "VirtualMachineScaleSetExtensionsClient.NewListPager":
		resp, err = v.dispatchNewListPager(req)
	case "VirtualMachineScaleSetExtensionsClient.BeginUpdate":
		resp, err = v.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (v *VirtualMachineScaleSetExtensionsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if v.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := v.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/virtualMachineScaleSets/(?P<vmScaleSetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/extensions/(?P<vmssExtensionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcompute.VirtualMachineScaleSetExtension](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		vmScaleSetNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("vmScaleSetName")])
		if err != nil {
			return nil, err
		}
		vmssExtensionNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("vmssExtensionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameUnescaped, vmScaleSetNameUnescaped, vmssExtensionNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		v.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		v.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		v.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (v *VirtualMachineScaleSetExtensionsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if v.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := v.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/virtualMachineScaleSets/(?P<vmScaleSetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/extensions/(?P<vmssExtensionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		vmScaleSetNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("vmScaleSetName")])
		if err != nil {
			return nil, err
		}
		vmssExtensionNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("vmssExtensionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginDelete(req.Context(), resourceGroupNameUnescaped, vmScaleSetNameUnescaped, vmssExtensionNameUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		v.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		v.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		v.beginDelete.remove(req)
	}

	return resp, nil
}

func (v *VirtualMachineScaleSetExtensionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if v.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/virtualMachineScaleSets/(?P<vmScaleSetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/extensions/(?P<vmssExtensionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	vmScaleSetNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("vmScaleSetName")])
	if err != nil {
		return nil, err
	}
	vmssExtensionNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("vmssExtensionName")])
	if err != nil {
		return nil, err
	}
	expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
	if err != nil {
		return nil, err
	}
	expandParam := getOptional(expandUnescaped)
	var options *armcompute.VirtualMachineScaleSetExtensionsClientGetOptions
	if expandParam != nil {
		options = &armcompute.VirtualMachineScaleSetExtensionsClientGetOptions{
			Expand: expandParam,
		}
	}
	respr, errRespr := v.srv.Get(req.Context(), resourceGroupNameUnescaped, vmScaleSetNameUnescaped, vmssExtensionNameUnescaped, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VirtualMachineScaleSetExtension, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VirtualMachineScaleSetExtensionsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if v.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := v.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/virtualMachineScaleSets/(?P<vmScaleSetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/extensions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		vmScaleSetNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("vmScaleSetName")])
		if err != nil {
			return nil, err
		}
		resp := v.srv.NewListPager(resourceGroupNameUnescaped, vmScaleSetNameUnescaped, nil)
		newListPager = &resp
		v.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armcompute.VirtualMachineScaleSetExtensionsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		v.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		v.newListPager.remove(req)
	}
	return resp, nil
}

func (v *VirtualMachineScaleSetExtensionsServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if v.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := v.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/virtualMachineScaleSets/(?P<vmScaleSetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/extensions/(?P<vmssExtensionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcompute.VirtualMachineScaleSetExtensionUpdate](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		vmScaleSetNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("vmScaleSetName")])
		if err != nil {
			return nil, err
		}
		vmssExtensionNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("vmssExtensionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginUpdate(req.Context(), resourceGroupNameUnescaped, vmScaleSetNameUnescaped, vmssExtensionNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		v.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		v.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		v.beginUpdate.remove(req)
	}

	return resp, nil
}
