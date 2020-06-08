<?php

// This file is auto-generated, don't edit it. Thanks.
namespace AlibabaCloud\SDK\Rds\V20140815\Models;

use AlibabaCloud\Tea\Model;

class CopyDatabaseResponse extends Model {
    protected $_name = [
        'DBName' => 'DBName',
        'DBStatus' => 'DBStatus',
        'taskId' => 'TaskId',
    ];
    public function validate() {
        Model::validateRequired('DBName', $this->DBName, true);
        Model::validateRequired('DBStatus', $this->DBStatus, true);
        Model::validateRequired('taskId', $this->taskId, true);
    }
    public function toMap() {
        $res = [];
        $res['DBName'] = $this->DBName;
        $res['DBStatus'] = $this->DBStatus;
        $res['TaskId'] = $this->taskId;
        return $res;
    }
    /**
     * @param array $map
     * @return CopyDatabaseResponse
     */
    public static function fromMap($map = []) {
        $model = new self();
        if(isset($map['DBName'])){
            $model->DBName = $map['DBName'];
        }
        if(isset($map['DBStatus'])){
            $model->DBStatus = $map['DBStatus'];
        }
        if(isset($map['TaskId'])){
            $model->taskId = $map['TaskId'];
        }
        return $model;
    }
    /**
     * @description dbName
     * @var string
     */
    public $DBName;

    /**
     * @description dbStatus
     * @var string
     */
    public $DBStatus;

    /**
     * @description taskId
     * @var string
     */
    public $taskId;

}