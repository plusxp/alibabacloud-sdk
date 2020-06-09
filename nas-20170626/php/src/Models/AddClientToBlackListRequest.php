<?php

// This file is auto-generated, don't edit it. Thanks.
namespace AlibabaCloud\SDK\NAS\V20170626\Models;

use AlibabaCloud\Tea\Model;

class AddClientToBlackListRequest extends Model {
    protected $_name = [
        'regionId' => 'RegionId',
        'fileSystemId' => 'FileSystemId',
        'clientIP' => 'ClientIP',
        'clientToken' => 'ClientToken',
    ];
    public function validate() {
        Model::validateRequired('regionId', $this->regionId, true);
        Model::validateRequired('fileSystemId', $this->fileSystemId, true);
        Model::validateRequired('clientIP', $this->clientIP, true);
        Model::validateRequired('clientToken', $this->clientToken, true);
    }
    public function toMap() {
        $res = [];
        $res['RegionId'] = $this->regionId;
        $res['FileSystemId'] = $this->fileSystemId;
        $res['ClientIP'] = $this->clientIP;
        $res['ClientToken'] = $this->clientToken;
        return $res;
    }
    /**
     * @param array $map
     * @return AddClientToBlackListRequest
     */
    public static function fromMap($map = []) {
        $model = new self();
        if(isset($map['RegionId'])){
            $model->regionId = $map['RegionId'];
        }
        if(isset($map['FileSystemId'])){
            $model->fileSystemId = $map['FileSystemId'];
        }
        if(isset($map['ClientIP'])){
            $model->clientIP = $map['ClientIP'];
        }
        if(isset($map['ClientToken'])){
            $model->clientToken = $map['ClientToken'];
        }
        return $model;
    }
    /**
     * @description region
     * @var string
     */
    public $regionId;

    /**
     * @description fileSystemId
     * @var string
     */
    public $fileSystemId;

    /**
     * @description clientIp
     * @var string
     */
    public $clientIP;

    /**
     * @description clientToken
     * @var string
     */
    public $clientToken;

}
