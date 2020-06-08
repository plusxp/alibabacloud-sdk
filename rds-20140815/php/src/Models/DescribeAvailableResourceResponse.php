<?php

// This file is auto-generated, don't edit it. Thanks.
namespace AlibabaCloud\SDK\Rds\V20140815\Models;

use AlibabaCloud\Tea\Model;

use AlibabaCloud\SDK\Rds\V20140815\Models\DescribeAvailableResourceResponse\availableZones;

class DescribeAvailableResourceResponse extends Model {
    protected $_name = [
        'requestId' => 'RequestId',
        'availableZones' => 'AvailableZones',
    ];
    public function validate() {
        Model::validateRequired('requestId', $this->requestId, true);
        Model::validateRequired('availableZones', $this->availableZones, true);
    }
    public function toMap() {
        $res = [];
        $res['RequestId'] = $this->requestId;
        $res['AvailableZones'] = null !== $this->availableZones ? $this->availableZones->toMap() : null;
        return $res;
    }
    /**
     * @param array $map
     * @return DescribeAvailableResourceResponse
     */
    public static function fromMap($map = []) {
        $model = new self();
        if(isset($map['RequestId'])){
            $model->requestId = $map['RequestId'];
        }
        if(isset($map['AvailableZones'])){
            $model->availableZones = availableZones::fromMap($map['AvailableZones']);
        }
        return $model;
    }
    /**
     * @description requestId
     * @var string
     */
    public $requestId;

    /**
     * @description data.availableZones
     * @var availableZones
     */
    public $availableZones;

}