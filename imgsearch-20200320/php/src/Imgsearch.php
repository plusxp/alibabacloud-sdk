<?php

// This file is auto-generated, don't edit it. Thanks.

namespace AlibabaCloud\SDK\Imgsearch\V20200320;

use AlibabaCloud\Endpoint\Endpoint;
use AlibabaCloud\SDK\Imgsearch\V20200320\Models\AddImageAdvanceRequest;
use AlibabaCloud\SDK\Imgsearch\V20200320\Models\AddImageRequest;
use AlibabaCloud\SDK\Imgsearch\V20200320\Models\AddImageResponse;
use AlibabaCloud\SDK\Imgsearch\V20200320\Models\CreateImageDbRequest;
use AlibabaCloud\SDK\Imgsearch\V20200320\Models\CreateImageDbResponse;
use AlibabaCloud\SDK\Imgsearch\V20200320\Models\DeleteImageDbRequest;
use AlibabaCloud\SDK\Imgsearch\V20200320\Models\DeleteImageDbResponse;
use AlibabaCloud\SDK\Imgsearch\V20200320\Models\DeleteImageRequest;
use AlibabaCloud\SDK\Imgsearch\V20200320\Models\DeleteImageResponse;
use AlibabaCloud\SDK\Imgsearch\V20200320\Models\ListImageDbsRequest;
use AlibabaCloud\SDK\Imgsearch\V20200320\Models\ListImageDbsResponse;
use AlibabaCloud\SDK\Imgsearch\V20200320\Models\ListImagesRequest;
use AlibabaCloud\SDK\Imgsearch\V20200320\Models\ListImagesResponse;
use AlibabaCloud\SDK\Imgsearch\V20200320\Models\SearchImageAdvanceRequest;
use AlibabaCloud\SDK\Imgsearch\V20200320\Models\SearchImageRequest;
use AlibabaCloud\SDK\Imgsearch\V20200320\Models\SearchImageResponse;
use AlibabaCloud\SDK\OpenPlatform\V20191219\Models\AuthorizeFileUploadRequest;
use AlibabaCloud\SDK\OpenPlatform\V20191219\OpenPlatform;
use AlibabaCloud\SDK\OSS\OSS;
use AlibabaCloud\SDK\OSS\OSS\PostObjectRequest;
use AlibabaCloud\SDK\OSS\OSS\PostObjectRequest\header;
use AlibabaCloud\Tea\FileForm\FileForm\FileField;
use AlibabaCloud\Tea\Rpc\Rpc;
use AlibabaCloud\Tea\Rpc\Rpc\Config;
use AlibabaCloud\Tea\RpcUtils\RpcUtils;
use AlibabaCloud\Tea\Utils\Utils;
use AlibabaCloud\Tea\Utils\Utils\RuntimeOptions;

class Imgsearch extends Rpc
{
    public function __construct($config)
    {
        parent::__construct($config);
        $this->_endpointRule = 'regional';
        $this->checkConfig($config);
        $this->_endpoint = $this->getEndpoint('imgsearch', $this->_regionId, $this->_endpointRule, $this->_network, $this->_suffix, $this->_endpointMap, $this->_endpoint);
    }

    /**
     * @param ListImageDbsRequest $request
     * @param RuntimeOptions      $runtime
     *
     * @return ListImageDbsResponse
     */
    public function listImageDbs($request, $runtime)
    {
        Utils::validateModel($request);

        return ListImageDbsResponse::fromMap($this->doRequest('ListImageDbs', 'HTTPS', 'POST', '2020-03-20', 'AK', null, $request, $runtime));
    }

    /**
     * @param ListImagesRequest $request
     * @param RuntimeOptions    $runtime
     *
     * @return ListImagesResponse
     */
    public function listImages($request, $runtime)
    {
        Utils::validateModel($request);

        return ListImagesResponse::fromMap($this->doRequest('ListImages', 'HTTPS', 'POST', '2020-03-20', 'AK', null, $request, $runtime));
    }

    /**
     * @param SearchImageRequest $request
     * @param RuntimeOptions     $runtime
     *
     * @return SearchImageResponse
     */
    public function searchImage($request, $runtime)
    {
        Utils::validateModel($request);

        return SearchImageResponse::fromMap($this->doRequest('SearchImage', 'HTTPS', 'POST', '2020-03-20', 'AK', null, $request, $runtime));
    }

    /**
     * @param SearchImageAdvanceRequest $request
     * @param RuntimeOptions            $runtime
     *
     * @return SearchImageResponse
     */
    public function searchImageAdvance($request, $runtime)
    {
        // Step 0: init client
        $accessKeyId     = $this->_credential->getAccessKeyId();
        $accessKeySecret = $this->_credential->getAccessKeySecret();
        $authConfig      = new Config([
            'accessKeyId'     => $accessKeyId,
            'accessKeySecret' => $accessKeySecret,
            'type'            => 'access_key',
            'endpoint'        => 'openplatform.aliyuncs.com',
            'protocol'        => $this->_protocol,
            'regionId'        => $this->_regionId,
        ]);
        $authClient  = new OpenPlatform($authConfig);
        $authRequest = new AuthorizeFileUploadRequest([
            'product'  => 'imgsearch',
            'regionId' => $this->_regionId,
        ]);
        $authResponse = $authClient->authorizeFileUploadWithOptions($authRequest, $runtime);
        // Step 1: request OSS api to upload file
        $ossConfig = new \AlibabaCloud\SDK\OSS\OSS\Config([
            'accessKeyId'     => $authResponse->accessKeyId,
            'accessKeySecret' => $accessKeySecret,
            'type'            => 'access_key',
            'endpoint'        => RpcUtils::getEndpoint($authResponse->endpoint, $authResponse->useAccelerate, $this->_endpointType),
            'protocol'        => $this->_protocol,
            'regionId'        => $this->_regionId,
        ]);
        $ossClient = new OSS($ossConfig);
        $fileObj   = new FileField([
            'filename'    => $authResponse->objectKey,
            'content'     => $request->imageUrlObject,
            'contentType' => '',
        ]);
        $ossHeader = new header([
            'accessKeyId'         => $authResponse->accessKeyId,
            'policy'              => $authResponse->encodedPolicy,
            'signature'           => $authResponse->signature,
            'key'                 => $authResponse->objectKey,
            'file'                => $fileObj,
            'successActionStatus' => '201',
        ]);
        $uploadRequest = new PostObjectRequest([
            'bucketName' => $authResponse->bucket,
            'header'     => $ossHeader,
        ]);
        $ossRuntime = new \AlibabaCloud\Tea\OSSUtils\OSSUtils\RuntimeOptions([]);
        RpcUtils::convert($runtime, $ossRuntime);
        $ossClient->postObject($uploadRequest, $ossRuntime);
        // Step 2: request final api
        $searchImagereq = new SearchImageRequest([]);
        RpcUtils::convert($request, $searchImagereq);
        $searchImagereq->imageUrl = 'http://' . $authResponse->bucket . '.' . $authResponse->endpoint . '/' . $authResponse->objectKey . '';

        return $this->searchImage($searchImagereq, $runtime);
    }

    /**
     * @param AddImageRequest $request
     * @param RuntimeOptions  $runtime
     *
     * @return AddImageResponse
     */
    public function addImage($request, $runtime)
    {
        Utils::validateModel($request);

        return AddImageResponse::fromMap($this->doRequest('AddImage', 'HTTPS', 'POST', '2020-03-20', 'AK', null, $request, $runtime));
    }

    /**
     * @param AddImageAdvanceRequest $request
     * @param RuntimeOptions         $runtime
     *
     * @return AddImageResponse
     */
    public function addImageAdvance($request, $runtime)
    {
        // Step 0: init client
        $accessKeyId     = $this->_credential->getAccessKeyId();
        $accessKeySecret = $this->_credential->getAccessKeySecret();
        $authConfig      = new Config([
            'accessKeyId'     => $accessKeyId,
            'accessKeySecret' => $accessKeySecret,
            'type'            => 'access_key',
            'endpoint'        => 'openplatform.aliyuncs.com',
            'protocol'        => $this->_protocol,
            'regionId'        => $this->_regionId,
        ]);
        $authClient  = new OpenPlatform($authConfig);
        $authRequest = new AuthorizeFileUploadRequest([
            'product'  => 'imgsearch',
            'regionId' => $this->_regionId,
        ]);
        $authResponse = $authClient->authorizeFileUploadWithOptions($authRequest, $runtime);
        // Step 1: request OSS api to upload file
        $ossConfig = new \AlibabaCloud\SDK\OSS\OSS\Config([
            'accessKeyId'     => $authResponse->accessKeyId,
            'accessKeySecret' => $accessKeySecret,
            'type'            => 'access_key',
            'endpoint'        => RpcUtils::getEndpoint($authResponse->endpoint, $authResponse->useAccelerate, $this->_endpointType),
            'protocol'        => $this->_protocol,
            'regionId'        => $this->_regionId,
        ]);
        $ossClient = new OSS($ossConfig);
        $fileObj   = new FileField([
            'filename'    => $authResponse->objectKey,
            'content'     => $request->imageUrlObject,
            'contentType' => '',
        ]);
        $ossHeader = new header([
            'accessKeyId'         => $authResponse->accessKeyId,
            'policy'              => $authResponse->encodedPolicy,
            'signature'           => $authResponse->signature,
            'key'                 => $authResponse->objectKey,
            'file'                => $fileObj,
            'successActionStatus' => '201',
        ]);
        $uploadRequest = new PostObjectRequest([
            'bucketName' => $authResponse->bucket,
            'header'     => $ossHeader,
        ]);
        $ossRuntime = new \AlibabaCloud\Tea\OSSUtils\OSSUtils\RuntimeOptions([]);
        RpcUtils::convert($runtime, $ossRuntime);
        $ossClient->postObject($uploadRequest, $ossRuntime);
        // Step 2: request final api
        $addImagereq = new AddImageRequest([]);
        RpcUtils::convert($request, $addImagereq);
        $addImagereq->imageUrl = 'http://' . $authResponse->bucket . '.' . $authResponse->endpoint . '/' . $authResponse->objectKey . '';

        return $this->addImage($addImagereq, $runtime);
    }

    /**
     * @param DeleteImageDbRequest $request
     * @param RuntimeOptions       $runtime
     *
     * @return DeleteImageDbResponse
     */
    public function deleteImageDb($request, $runtime)
    {
        Utils::validateModel($request);

        return DeleteImageDbResponse::fromMap($this->doRequest('DeleteImageDb', 'HTTPS', 'POST', '2020-03-20', 'AK', null, $request, $runtime));
    }

    /**
     * @param CreateImageDbRequest $request
     * @param RuntimeOptions       $runtime
     *
     * @return CreateImageDbResponse
     */
    public function createImageDb($request, $runtime)
    {
        Utils::validateModel($request);

        return CreateImageDbResponse::fromMap($this->doRequest('CreateImageDb', 'HTTPS', 'POST', '2020-03-20', 'AK', null, $request, $runtime));
    }

    /**
     * @param DeleteImageRequest $request
     * @param RuntimeOptions     $runtime
     *
     * @return DeleteImageResponse
     */
    public function deleteImage($request, $runtime)
    {
        Utils::validateModel($request);

        return DeleteImageResponse::fromMap($this->doRequest('DeleteImage', 'HTTPS', 'POST', '2020-03-20', 'AK', null, $request, $runtime));
    }

    /**
     * @param string $productId
     * @param string $regionId
     * @param string $endpointRule
     * @param string $network
     * @param string $suffix
     * @param array  $endpointMap
     * @param string $endpoint
     *
     * @return string
     */
    public function getEndpoint($productId, $regionId, $endpointRule, $network, $suffix, $endpointMap, $endpoint)
    {
        if (!Utils::empty_($endpoint)) {
            return $endpoint;
        }
        if (!Utils::isUnset($endpointMap) && !Utils::empty_($endpointMap['regionId'])) {
            return $endpointMap['regionId'];
        }

        return Endpoint::getEndpointRules($productId, $regionId, $endpointRule, $network, $suffix);
    }
}
