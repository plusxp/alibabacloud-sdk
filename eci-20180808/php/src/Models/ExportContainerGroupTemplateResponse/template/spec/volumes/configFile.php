<?php

// This file is auto-generated, don't edit it. Thanks.

namespace AlibabaCloud\SDK\Eci\V20180808\Models\ExportContainerGroupTemplateResponse\template\spec\volumes;

use AlibabaCloud\SDK\Eci\V20180808\Models\ExportContainerGroupTemplateResponse\template\spec\volumes\configFile\items;
use AlibabaCloud\Tea\Model;

class configFile extends Model
{
    /**
     * @description defaultMode
     *
     * @var int
     */
    public $defaultMode;

    /**
     * @description items
     *
     * @var array
     */
    public $items;
    protected $_name = [
        'defaultMode' => 'DefaultMode',
        'items'       => 'Items',
    ];

    public function validate()
    {
        Model::validateRequired('defaultMode', $this->defaultMode, true);
        Model::validateRequired('items', $this->items, true);
    }

    public function toMap()
    {
        $res = [];
        if (null !== $this->defaultMode) {
            $res['DefaultMode'] = $this->defaultMode;
        }
        if (null !== $this->items) {
            $res['Items'] = [];
            if (null !== $this->items && \is_array($this->items)) {
                $n = 0;
                foreach ($this->items as $item) {
                    $res['Items'][$n++] = null !== $item ? $item->toMap() : $item;
                }
            }
        }

        return $res;
    }

    /**
     * @param array $map
     *
     * @return configFile
     */
    public static function fromMap($map = [])
    {
        $model = new self();
        if (isset($map['DefaultMode'])) {
            $model->defaultMode = $map['DefaultMode'];
        }
        if (isset($map['Items'])) {
            if (!empty($map['Items'])) {
                $model->items = [];
                $n            = 0;
                foreach ($map['Items'] as $item) {
                    $model->items[$n++] = null !== $item ? items::fromMap($item) : $item;
                }
            }
        }

        return $model;
    }
}
