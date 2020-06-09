<?php

// This file is auto-generated, don't edit it. Thanks.
namespace AlibabaCloud\SDK\NAS\V20170626\Models;

use AlibabaCloud\Tea\Model;

class CreateFileSystemResponse extends Model {
    protected $_name = [
        'requestId' => 'RequestId',
        'fileSystemId' => 'FileSystemId',
    ];
    public function validate() {
        Model::validateRequired('requestId', $this->requestId, true);
        Model::validateRequired('fileSystemId', $this->fileSystemId, true);
    }
    public function toMap() {
        $res = [];
        $res['RequestId'] = $this->requestId;
        $res['FileSystemId'] = $this->fileSystemId;
        return $res;
    }
    /**
     * @param array $map
     * @return CreateFileSystemResponse
     */
    public static function fromMap($map = []) {
        $model = new self();
        if(isset($map['RequestId'])){
            $model->requestId = $map['RequestId'];
        }
        if(isset($map['FileSystemId'])){
            $model->fileSystemId = $map['FileSystemId'];
        }
        return $model;
    }
    /**
     * @description requestId
     * @var string
     */
    public $requestId;

    /**
     * @description data
     * @var string
     */
    public $fileSystemId;

}