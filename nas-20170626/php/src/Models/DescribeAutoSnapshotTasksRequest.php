<?php

// This file is auto-generated, don't edit it. Thanks.
namespace AlibabaCloud\SDK\NAS\V20170626\Models;

use AlibabaCloud\Tea\Model;

class DescribeAutoSnapshotTasksRequest extends Model {
    protected $_name = [
        'fileSystemIds' => 'FileSystemIds',
        'autoSnapshotPolicyIds' => 'AutoSnapshotPolicyIds',
        'fileSystemType' => 'FileSystemType',
        'pageSize' => 'PageSize',
        'pageNumber' => 'PageNumber',
    ];
    public function validate() {
        Model::validateRequired('fileSystemType', $this->fileSystemType, true);
    }
    public function toMap() {
        $res = [];
        $res['FileSystemIds'] = $this->fileSystemIds;
        $res['AutoSnapshotPolicyIds'] = $this->autoSnapshotPolicyIds;
        $res['FileSystemType'] = $this->fileSystemType;
        $res['PageSize'] = $this->pageSize;
        $res['PageNumber'] = $this->pageNumber;
        return $res;
    }
    /**
     * @param array $map
     * @return DescribeAutoSnapshotTasksRequest
     */
    public static function fromMap($map = []) {
        $model = new self();
        if(isset($map['FileSystemIds'])){
            $model->fileSystemIds = $map['FileSystemIds'];
        }
        if(isset($map['AutoSnapshotPolicyIds'])){
            $model->autoSnapshotPolicyIds = $map['AutoSnapshotPolicyIds'];
        }
        if(isset($map['FileSystemType'])){
            $model->fileSystemType = $map['FileSystemType'];
        }
        if(isset($map['PageSize'])){
            $model->pageSize = $map['PageSize'];
        }
        if(isset($map['PageNumber'])){
            $model->pageNumber = $map['PageNumber'];
        }
        return $model;
    }
    /**
     * @description volumes
     * @var string
     */
    public $fileSystemIds;

    /**
     * @description autoSnapshotPolicyIds
     * @var string
     */
    public $autoSnapshotPolicyIds;

    /**
     * @description fileSystemType
     * @var string
     */
    public $fileSystemType;

    /**
     * @description pageSize
     * @var integer
     */
    public $pageSize;

    /**
     * @description pageNumber
     * @var integer
     */
    public $pageNumber;

}