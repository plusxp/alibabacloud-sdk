<?php

// This file is auto-generated, don't edit it. Thanks.
namespace AlibabaCloud\SDK\NAS\V20170626\Models;

use AlibabaCloud\Tea\Model;

class UpgradeFileSystemRequest extends Model {
    protected $_name = [
        'fileSystemId' => 'FileSystemId',
        'capacity' => 'Capacity',
        'dryRun' => 'DryRun',
        'clientToken' => 'ClientToken',
    ];
    public function validate() {
        Model::validateRequired('fileSystemId', $this->fileSystemId, true);
        Model::validateRequired('capacity', $this->capacity, true);
    }
    public function toMap() {
        $res = [];
        $res['FileSystemId'] = $this->fileSystemId;
        $res['Capacity'] = $this->capacity;
        $res['DryRun'] = $this->dryRun;
        $res['ClientToken'] = $this->clientToken;
        return $res;
    }
    /**
     * @param array $map
     * @return UpgradeFileSystemRequest
     */
    public static function fromMap($map = []) {
        $model = new self();
        if(isset($map['FileSystemId'])){
            $model->fileSystemId = $map['FileSystemId'];
        }
        if(isset($map['Capacity'])){
            $model->capacity = $map['Capacity'];
        }
        if(isset($map['DryRun'])){
            $model->dryRun = $map['DryRun'];
        }
        if(isset($map['ClientToken'])){
            $model->clientToken = $map['ClientToken'];
        }
        return $model;
    }
    /**
     * @description volume
     * @var string
     */
    public $fileSystemId;

    /**
     * @description capacity
     * @var integer
     */
    public $capacity;

    /**
     * @description dryRun
     * @var bool
     */
    public $dryRun;

    /**
     * @description clientToken
     * @var string
     */
    public $clientToken;

}
