<?php

// This file is auto-generated, don't edit it. Thanks.

namespace AlibabaCloud\SDK\Ecs\V20140526\Models;

use AlibabaCloud\Tea\Model;

class GetInstanceScreenshotRequest extends Model
{
    /**
     * @var string
     */
    public $regionId;

    /**
     * @var string
     */
    public $instanceId;

    /**
     * @var bool
     */
    public $wakeUp;
    protected $_name = [
        'regionId'   => 'RegionId',
        'instanceId' => 'InstanceId',
        'wakeUp'     => 'WakeUp',
    ];

    public function validate()
    {
        Model::validateRequired('regionId', $this->regionId, true);
        Model::validateRequired('instanceId', $this->instanceId, true);
    }

    public function toMap()
    {
        $res = [];
        if (null !== $this->regionId) {
            $res['RegionId'] = $this->regionId;
        }
        if (null !== $this->instanceId) {
            $res['InstanceId'] = $this->instanceId;
        }
        if (null !== $this->wakeUp) {
            $res['WakeUp'] = $this->wakeUp;
        }

        return $res;
    }

    /**
     * @param array $map
     *
     * @return GetInstanceScreenshotRequest
     */
    public static function fromMap($map = [])
    {
        $model = new self();
        if (isset($map['RegionId'])) {
            $model->regionId = $map['RegionId'];
        }
        if (isset($map['InstanceId'])) {
            $model->instanceId = $map['InstanceId'];
        }
        if (isset($map['WakeUp'])) {
            $model->wakeUp = $map['WakeUp'];
        }

        return $model;
    }
}
