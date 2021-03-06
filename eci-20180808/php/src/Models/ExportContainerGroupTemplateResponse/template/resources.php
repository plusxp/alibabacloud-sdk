<?php

// This file is auto-generated, don't edit it. Thanks.

namespace AlibabaCloud\SDK\Eci\V20180808\Models\ExportContainerGroupTemplateResponse\template;

use AlibabaCloud\Tea\Model;

class resources extends Model
{
    /**
     * @description bigCpuForPop
     *
     * @var float
     */
    public $cpu;

    /**
     * @description bigMemForPop
     *
     * @var float
     */
    public $memory;
    protected $_name = [
        'cpu'    => 'Cpu',
        'memory' => 'Memory',
    ];

    public function validate()
    {
        Model::validateRequired('cpu', $this->cpu, true);
        Model::validateRequired('memory', $this->memory, true);
    }

    public function toMap()
    {
        $res = [];
        if (null !== $this->cpu) {
            $res['Cpu'] = $this->cpu;
        }
        if (null !== $this->memory) {
            $res['Memory'] = $this->memory;
        }

        return $res;
    }

    /**
     * @param array $map
     *
     * @return resources
     */
    public static function fromMap($map = [])
    {
        $model = new self();
        if (isset($map['Cpu'])) {
            $model->cpu = $map['Cpu'];
        }
        if (isset($map['Memory'])) {
            $model->memory = $map['Memory'];
        }

        return $model;
    }
}
