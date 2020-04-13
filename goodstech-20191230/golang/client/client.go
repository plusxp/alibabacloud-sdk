// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/tea"
	oss "github.com/aliyun/alibabacloud-oss-sdk/golang/client"
	ossutil "github.com/aliyun/alibabacloud-oss-sdk/util/golang/service"
	rpcutil "github.com/aliyun/alibabacloud-rpc-util-sdk/golang/service"
	openplatform "github.com/aliyun/alibabacloud-sdk/openplatform-20191219/golang/client"
	endpointutil "github.com/aliyun/endpoint-util/golang/service"
	fileform "github.com/aliyun/tea-fileform/golang/service"
	rpc "github.com/aliyun/tea-rpc/golang/client"
	util "github.com/aliyun/tea-util/golang/service"
	"io"
)

type RecognizeFurnitureAttributeRequest struct {
	ImageURL *string `json:"ImageURL" xml:"ImageURL" require:"true"`
}

func (s RecognizeFurnitureAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFurnitureAttributeRequest) GoString() string {
	return s.String()
}

func (s *RecognizeFurnitureAttributeRequest) SetImageURL(v string) *RecognizeFurnitureAttributeRequest {
	s.ImageURL = &v
	return s
}

type RecognizeFurnitureAttributeResponse struct {
	RequestId *string                                  `json:"RequestId" xml:"RequestId" require:"true"`
	Data      *RecognizeFurnitureAttributeResponseData `json:"Data" xml:"Data" require:"true" type:"Struct"`
}

func (s RecognizeFurnitureAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFurnitureAttributeResponse) GoString() string {
	return s.String()
}

func (s *RecognizeFurnitureAttributeResponse) SetRequestId(v string) *RecognizeFurnitureAttributeResponse {
	s.RequestId = &v
	return s
}

func (s *RecognizeFurnitureAttributeResponse) SetData(v *RecognizeFurnitureAttributeResponseData) *RecognizeFurnitureAttributeResponse {
	s.Data = v
	return s
}

type RecognizeFurnitureAttributeResponseData struct {
	PredStyleId     *string  `json:"PredStyleId" xml:"PredStyleId" require:"true"`
	PredStyle       *string  `json:"PredStyle" xml:"PredStyle" require:"true"`
	PredProbability *float32 `json:"PredProbability" xml:"PredProbability" require:"true"`
}

func (s RecognizeFurnitureAttributeResponseData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFurnitureAttributeResponseData) GoString() string {
	return s.String()
}

func (s *RecognizeFurnitureAttributeResponseData) SetPredStyleId(v string) *RecognizeFurnitureAttributeResponseData {
	s.PredStyleId = &v
	return s
}

func (s *RecognizeFurnitureAttributeResponseData) SetPredStyle(v string) *RecognizeFurnitureAttributeResponseData {
	s.PredStyle = &v
	return s
}

func (s *RecognizeFurnitureAttributeResponseData) SetPredProbability(v float32) *RecognizeFurnitureAttributeResponseData {
	s.PredProbability = &v
	return s
}

type RecognizeFurnitureAttributeAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject" xml:"ImageURLObject" require:"true"`
}

func (s RecognizeFurnitureAttributeAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFurnitureAttributeAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeFurnitureAttributeAdvanceRequest) SetImageURLObject(v io.Reader) *RecognizeFurnitureAttributeAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type RecognizeFurnitureSpuRequest struct {
	ImageURL *string  `json:"ImageURL" xml:"ImageURL" require:"true"`
	XLength  *float32 `json:"XLength" xml:"XLength" require:"true"`
	YLength  *float32 `json:"YLength" xml:"YLength" require:"true"`
	ZLength  *float32 `json:"ZLength" xml:"ZLength" require:"true"`
}

func (s RecognizeFurnitureSpuRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFurnitureSpuRequest) GoString() string {
	return s.String()
}

func (s *RecognizeFurnitureSpuRequest) SetImageURL(v string) *RecognizeFurnitureSpuRequest {
	s.ImageURL = &v
	return s
}

func (s *RecognizeFurnitureSpuRequest) SetXLength(v float32) *RecognizeFurnitureSpuRequest {
	s.XLength = &v
	return s
}

func (s *RecognizeFurnitureSpuRequest) SetYLength(v float32) *RecognizeFurnitureSpuRequest {
	s.YLength = &v
	return s
}

func (s *RecognizeFurnitureSpuRequest) SetZLength(v float32) *RecognizeFurnitureSpuRequest {
	s.ZLength = &v
	return s
}

type RecognizeFurnitureSpuResponse struct {
	RequestId *string                            `json:"RequestId" xml:"RequestId" require:"true"`
	Data      *RecognizeFurnitureSpuResponseData `json:"Data" xml:"Data" require:"true" type:"Struct"`
}

func (s RecognizeFurnitureSpuResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFurnitureSpuResponse) GoString() string {
	return s.String()
}

func (s *RecognizeFurnitureSpuResponse) SetRequestId(v string) *RecognizeFurnitureSpuResponse {
	s.RequestId = &v
	return s
}

func (s *RecognizeFurnitureSpuResponse) SetData(v *RecognizeFurnitureSpuResponseData) *RecognizeFurnitureSpuResponse {
	s.Data = v
	return s
}

type RecognizeFurnitureSpuResponseData struct {
	PredCateId      *string  `json:"PredCateId" xml:"PredCateId" require:"true"`
	PredCate        *string  `json:"PredCate" xml:"PredCate" require:"true"`
	PredProbability *float32 `json:"PredProbability" xml:"PredProbability" require:"true"`
}

func (s RecognizeFurnitureSpuResponseData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFurnitureSpuResponseData) GoString() string {
	return s.String()
}

func (s *RecognizeFurnitureSpuResponseData) SetPredCateId(v string) *RecognizeFurnitureSpuResponseData {
	s.PredCateId = &v
	return s
}

func (s *RecognizeFurnitureSpuResponseData) SetPredCate(v string) *RecognizeFurnitureSpuResponseData {
	s.PredCate = &v
	return s
}

func (s *RecognizeFurnitureSpuResponseData) SetPredProbability(v float32) *RecognizeFurnitureSpuResponseData {
	s.PredProbability = &v
	return s
}

type RecognizeFurnitureSpuAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject" xml:"ImageURLObject" require:"true"`
	XLength        *float32  `json:"XLength" xml:"XLength" require:"true"`
	YLength        *float32  `json:"YLength" xml:"YLength" require:"true"`
	ZLength        *float32  `json:"ZLength" xml:"ZLength" require:"true"`
}

func (s RecognizeFurnitureSpuAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFurnitureSpuAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeFurnitureSpuAdvanceRequest) SetImageURLObject(v io.Reader) *RecognizeFurnitureSpuAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *RecognizeFurnitureSpuAdvanceRequest) SetXLength(v float32) *RecognizeFurnitureSpuAdvanceRequest {
	s.XLength = &v
	return s
}

func (s *RecognizeFurnitureSpuAdvanceRequest) SetYLength(v float32) *RecognizeFurnitureSpuAdvanceRequest {
	s.YLength = &v
	return s
}

func (s *RecognizeFurnitureSpuAdvanceRequest) SetZLength(v float32) *RecognizeFurnitureSpuAdvanceRequest {
	s.ZLength = &v
	return s
}

type ClassifyCommodityRequest struct {
	ImageURL *string `json:"ImageURL" xml:"ImageURL" require:"true"`
}

func (s ClassifyCommodityRequest) String() string {
	return tea.Prettify(s)
}

func (s ClassifyCommodityRequest) GoString() string {
	return s.String()
}

func (s *ClassifyCommodityRequest) SetImageURL(v string) *ClassifyCommodityRequest {
	s.ImageURL = &v
	return s
}

type ClassifyCommodityResponse struct {
	RequestId *string                        `json:"RequestId" xml:"RequestId" require:"true"`
	Data      *ClassifyCommodityResponseData `json:"Data" xml:"Data" require:"true" type:"Struct"`
}

func (s ClassifyCommodityResponse) String() string {
	return tea.Prettify(s)
}

func (s ClassifyCommodityResponse) GoString() string {
	return s.String()
}

func (s *ClassifyCommodityResponse) SetRequestId(v string) *ClassifyCommodityResponse {
	s.RequestId = &v
	return s
}

func (s *ClassifyCommodityResponse) SetData(v *ClassifyCommodityResponseData) *ClassifyCommodityResponse {
	s.Data = v
	return s
}

type ClassifyCommodityResponseData struct {
	Categories []*ClassifyCommodityResponseDataCategories `json:"Categories" xml:"Categories" require:"true" type:"Repeated"`
}

func (s ClassifyCommodityResponseData) String() string {
	return tea.Prettify(s)
}

func (s ClassifyCommodityResponseData) GoString() string {
	return s.String()
}

func (s *ClassifyCommodityResponseData) SetCategories(v []*ClassifyCommodityResponseDataCategories) *ClassifyCommodityResponseData {
	s.Categories = v
	return s
}

type ClassifyCommodityResponseDataCategories struct {
	Score        *float32 `json:"Score" xml:"Score" require:"true"`
	CategoryName *string  `json:"CategoryName" xml:"CategoryName" require:"true"`
	CategoryId   *string  `json:"CategoryId" xml:"CategoryId" require:"true"`
}

func (s ClassifyCommodityResponseDataCategories) String() string {
	return tea.Prettify(s)
}

func (s ClassifyCommodityResponseDataCategories) GoString() string {
	return s.String()
}

func (s *ClassifyCommodityResponseDataCategories) SetScore(v float32) *ClassifyCommodityResponseDataCategories {
	s.Score = &v
	return s
}

func (s *ClassifyCommodityResponseDataCategories) SetCategoryName(v string) *ClassifyCommodityResponseDataCategories {
	s.CategoryName = &v
	return s
}

func (s *ClassifyCommodityResponseDataCategories) SetCategoryId(v string) *ClassifyCommodityResponseDataCategories {
	s.CategoryId = &v
	return s
}

type ClassifyCommodityAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject" xml:"ImageURLObject" require:"true"`
}

func (s ClassifyCommodityAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ClassifyCommodityAdvanceRequest) GoString() string {
	return s.String()
}

func (s *ClassifyCommodityAdvanceRequest) SetImageURLObject(v io.Reader) *ClassifyCommodityAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type Client struct {
	rpc.Client
}

func NewClient(config *rpc.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *rpc.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = "regional"
	_err = client.CheckConfig(config)
	if _err != nil {
		return
	}
	client.Endpoint, _err = client.GetEndpoint(client.ProductId, client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) RecognizeFurnitureAttribute(request *RecognizeFurnitureAttributeRequest, runtime *util.RuntimeOptions) (_result *RecognizeFurnitureAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return
	}
	_result = &RecognizeFurnitureAttributeResponse{}
	_body, _err := client.DoRequest("RecognizeFurnitureAttribute", "HTTPS", "POST", "2019-12-30", "AK", nil, tea.ToMap(request), runtime)
	if _err != nil {
		return nil, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeFurnitureAttributeAdvance(request *RecognizeFurnitureAttributeAdvanceRequest, runtime *util.RuntimeOptions) (_result *RecognizeFurnitureAttributeResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return nil, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return nil, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     tea.String(accessKeyId),
		AccessKeySecret: tea.String(accessKeySecret),
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        tea.String(client.Protocol),
		RegionId:        tea.String(client.RegionId),
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return nil, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("goodstech"),
		RegionId: tea.String(client.RegionId),
	}
	authResponse, _err := authClient.AuthorizeFileUpload(authRequest, runtime)
	if _err != nil {
		return nil, _err
	}

	// Step 1: request OSS api to upload file
	ossConfig := &oss.Config{
		AccessKeyId:     authResponse.AccessKeyId,
		AccessKeySecret: tea.String(accessKeySecret),
		Type:            tea.String("access_key"),
		Endpoint:        tea.String(rpcutil.GetEndpoint(tea.StringValue(authResponse.Endpoint), tea.BoolValue(authResponse.UseAccelerate), client.EndpointType)),
		Protocol:        tea.String(client.Protocol),
		RegionId:        tea.String(client.RegionId),
	}
	ossClient, _err := oss.NewClient(ossConfig)
	if _err != nil {
		return nil, _err
	}

	fileObj := &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader := &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest := &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	ossRuntime := &ossutil.RuntimeOptions{}
	rpcutil.Convert(runtime, ossRuntime)
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return
	}
	// Step 2: request final api
	recognizeFurnitureAttributereq := &RecognizeFurnitureAttributeRequest{}
	rpcutil.Convert(request, recognizeFurnitureAttributereq)
	recognizeFurnitureAttributereq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	recognizeFurnitureAttributeResp, _err := client.RecognizeFurnitureAttribute(recognizeFurnitureAttributereq, runtime)
	if _err != nil {
		return nil, _err
	}

	_result = recognizeFurnitureAttributeResp
	return _result, _err
}

func (client *Client) RecognizeFurnitureSpu(request *RecognizeFurnitureSpuRequest, runtime *util.RuntimeOptions) (_result *RecognizeFurnitureSpuResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return
	}
	_result = &RecognizeFurnitureSpuResponse{}
	_body, _err := client.DoRequest("RecognizeFurnitureSpu", "HTTPS", "POST", "2019-12-30", "AK", nil, tea.ToMap(request), runtime)
	if _err != nil {
		return nil, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeFurnitureSpuAdvance(request *RecognizeFurnitureSpuAdvanceRequest, runtime *util.RuntimeOptions) (_result *RecognizeFurnitureSpuResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return nil, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return nil, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     tea.String(accessKeyId),
		AccessKeySecret: tea.String(accessKeySecret),
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        tea.String(client.Protocol),
		RegionId:        tea.String(client.RegionId),
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return nil, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("goodstech"),
		RegionId: tea.String(client.RegionId),
	}
	authResponse, _err := authClient.AuthorizeFileUpload(authRequest, runtime)
	if _err != nil {
		return nil, _err
	}

	// Step 1: request OSS api to upload file
	ossConfig := &oss.Config{
		AccessKeyId:     authResponse.AccessKeyId,
		AccessKeySecret: tea.String(accessKeySecret),
		Type:            tea.String("access_key"),
		Endpoint:        tea.String(rpcutil.GetEndpoint(tea.StringValue(authResponse.Endpoint), tea.BoolValue(authResponse.UseAccelerate), client.EndpointType)),
		Protocol:        tea.String(client.Protocol),
		RegionId:        tea.String(client.RegionId),
	}
	ossClient, _err := oss.NewClient(ossConfig)
	if _err != nil {
		return nil, _err
	}

	fileObj := &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader := &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest := &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	ossRuntime := &ossutil.RuntimeOptions{}
	rpcutil.Convert(runtime, ossRuntime)
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return
	}
	// Step 2: request final api
	recognizeFurnitureSpureq := &RecognizeFurnitureSpuRequest{}
	rpcutil.Convert(request, recognizeFurnitureSpureq)
	recognizeFurnitureSpureq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	recognizeFurnitureSpuResp, _err := client.RecognizeFurnitureSpu(recognizeFurnitureSpureq, runtime)
	if _err != nil {
		return nil, _err
	}

	_result = recognizeFurnitureSpuResp
	return _result, _err
}

func (client *Client) ClassifyCommodity(request *ClassifyCommodityRequest, runtime *util.RuntimeOptions) (_result *ClassifyCommodityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return
	}
	_result = &ClassifyCommodityResponse{}
	_body, _err := client.DoRequest("ClassifyCommodity", "HTTPS", "GET", "2019-12-30", "AK", tea.ToMap(request), nil, runtime)
	if _err != nil {
		return nil, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ClassifyCommodityAdvance(request *ClassifyCommodityAdvanceRequest, runtime *util.RuntimeOptions) (_result *ClassifyCommodityResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return nil, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return nil, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     tea.String(accessKeyId),
		AccessKeySecret: tea.String(accessKeySecret),
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        tea.String(client.Protocol),
		RegionId:        tea.String(client.RegionId),
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return nil, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("goodstech"),
		RegionId: tea.String(client.RegionId),
	}
	authResponse, _err := authClient.AuthorizeFileUpload(authRequest, runtime)
	if _err != nil {
		return nil, _err
	}

	// Step 1: request OSS api to upload file
	ossConfig := &oss.Config{
		AccessKeyId:     authResponse.AccessKeyId,
		AccessKeySecret: tea.String(accessKeySecret),
		Type:            tea.String("access_key"),
		Endpoint:        tea.String(rpcutil.GetEndpoint(tea.StringValue(authResponse.Endpoint), tea.BoolValue(authResponse.UseAccelerate), client.EndpointType)),
		Protocol:        tea.String(client.Protocol),
		RegionId:        tea.String(client.RegionId),
	}
	ossClient, _err := oss.NewClient(ossConfig)
	if _err != nil {
		return nil, _err
	}

	fileObj := &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader := &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest := &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	ossRuntime := &ossutil.RuntimeOptions{}
	rpcutil.Convert(runtime, ossRuntime)
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return
	}
	// Step 2: request final api
	classifyCommodityreq := &ClassifyCommodityRequest{}
	rpcutil.Convert(request, classifyCommodityreq)
	classifyCommodityreq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	classifyCommodityResp, _err := client.ClassifyCommodity(classifyCommodityreq, runtime)
	if _err != nil {
		return nil, _err
	}

	_result = classifyCommodityResp
	return _result, _err
}

func (client *Client) GetEndpoint(productId string, regionId string, endpointRule string, network string, suffix string, endpointMap map[string]string, endpoint string) (_result string, _err error) {
	if !util.Empty(endpoint) {
		_result = endpoint
		return _result, _err
	}

	if !util.IsUnset(endpointMap) && !util.Empty(endpointMap[regionId]) {
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return "", _err
	}
	_result = _body
	return _result, _err
}
