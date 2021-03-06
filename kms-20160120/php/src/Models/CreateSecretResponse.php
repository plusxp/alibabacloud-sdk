<?php

// This file is auto-generated, don't edit it. Thanks.

namespace AlibabaCloud\SDK\Kms\V20160120\Models;

use AlibabaCloud\Tea\Model;

class CreateSecretResponse extends Model
{
    /**
     * @description RequestId
     *
     * @var string
     */
    public $requestId;

    /**
     * @description Arn
     *
     * @var string
     */
    public $arn;

    /**
     * @description VersionId
     *
     * @var string
     */
    public $versionId;

    /**
     * @description SecretName
     *
     * @var string
     */
    public $secretName;
    protected $_name = [
        'requestId'  => 'RequestId',
        'arn'        => 'Arn',
        'versionId'  => 'VersionId',
        'secretName' => 'SecretName',
    ];

    public function validate()
    {
        Model::validateRequired('requestId', $this->requestId, true);
        Model::validateRequired('arn', $this->arn, true);
        Model::validateRequired('versionId', $this->versionId, true);
        Model::validateRequired('secretName', $this->secretName, true);
    }

    public function toMap()
    {
        $res = [];
        if (null !== $this->requestId) {
            $res['RequestId'] = $this->requestId;
        }
        if (null !== $this->arn) {
            $res['Arn'] = $this->arn;
        }
        if (null !== $this->versionId) {
            $res['VersionId'] = $this->versionId;
        }
        if (null !== $this->secretName) {
            $res['SecretName'] = $this->secretName;
        }

        return $res;
    }

    /**
     * @param array $map
     *
     * @return CreateSecretResponse
     */
    public static function fromMap($map = [])
    {
        $model = new self();
        if (isset($map['RequestId'])) {
            $model->requestId = $map['RequestId'];
        }
        if (isset($map['Arn'])) {
            $model->arn = $map['Arn'];
        }
        if (isset($map['VersionId'])) {
            $model->versionId = $map['VersionId'];
        }
        if (isset($map['SecretName'])) {
            $model->secretName = $map['SecretName'];
        }

        return $model;
    }
}
