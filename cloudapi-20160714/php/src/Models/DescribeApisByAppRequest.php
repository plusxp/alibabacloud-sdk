<?php

// This file is auto-generated, don't edit it. Thanks.

namespace AlibabaCloud\SDK\CloudAPI\V20160714\Models;

use AlibabaCloud\Tea\Model;

class DescribeApisByAppRequest extends Model
{
    /**
     * @description securityToken
     *
     * @var string
     */
    public $securityToken;

    /**
     * @description accessKeyId
     *
     * @var string
     */
    public $accessKeyId;

    /**
     * @description appId
     *
     * @var int
     */
    public $appId;

    /**
     * @description pageNumber
     *
     * @var int
     */
    public $pageNumber;

    /**
     * @description pageSize
     *
     * @var int
     */
    public $pageSize;

    /**
     * @description apiUid
     *
     * @var string
     */
    public $apiUid;
    protected $_name = [
        'securityToken' => 'SecurityToken',
        'accessKeyId'   => 'AccessKeyId',
        'appId'         => 'AppId',
        'pageNumber'    => 'PageNumber',
        'pageSize'      => 'PageSize',
        'apiUid'        => 'ApiUid',
    ];

    public function validate()
    {
        Model::validateRequired('appId', $this->appId, true);
    }

    public function toMap()
    {
        $res = [];
        if (null !== $this->securityToken) {
            $res['SecurityToken'] = $this->securityToken;
        }
        if (null !== $this->accessKeyId) {
            $res['AccessKeyId'] = $this->accessKeyId;
        }
        if (null !== $this->appId) {
            $res['AppId'] = $this->appId;
        }
        if (null !== $this->pageNumber) {
            $res['PageNumber'] = $this->pageNumber;
        }
        if (null !== $this->pageSize) {
            $res['PageSize'] = $this->pageSize;
        }
        if (null !== $this->apiUid) {
            $res['ApiUid'] = $this->apiUid;
        }

        return $res;
    }

    /**
     * @param array $map
     *
     * @return DescribeApisByAppRequest
     */
    public static function fromMap($map = [])
    {
        $model = new self();
        if (isset($map['SecurityToken'])) {
            $model->securityToken = $map['SecurityToken'];
        }
        if (isset($map['AccessKeyId'])) {
            $model->accessKeyId = $map['AccessKeyId'];
        }
        if (isset($map['AppId'])) {
            $model->appId = $map['AppId'];
        }
        if (isset($map['PageNumber'])) {
            $model->pageNumber = $map['PageNumber'];
        }
        if (isset($map['PageSize'])) {
            $model->pageSize = $map['PageSize'];
        }
        if (isset($map['ApiUid'])) {
            $model->apiUid = $map['ApiUid'];
        }

        return $model;
    }
}
