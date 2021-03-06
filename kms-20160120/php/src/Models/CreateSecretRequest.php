<?php

// This file is auto-generated, don't edit it. Thanks.

namespace AlibabaCloud\SDK\Kms\V20160120\Models;

use AlibabaCloud\Tea\Model;

class CreateSecretRequest extends Model
{
    /**
     * @description SecretName
     *
     * @var string
     */
    public $secretName;

    /**
     * @description VersionId
     *
     * @var string
     */
    public $versionId;

    /**
     * @description EncryptionKeyId
     *
     * @var string
     */
    public $encryptionKeyId;

    /**
     * @description SecretData
     *
     * @var string
     */
    public $secretData;

    /**
     * @description SecretDataType
     *
     * @var string
     */
    public $secretDataType;

    /**
     * @description Description
     *
     * @var string
     */
    public $description;

    /**
     * @description Tags
     *
     * @var string
     */
    public $tags;
    protected $_name = [
        'secretName'      => 'SecretName',
        'versionId'       => 'VersionId',
        'encryptionKeyId' => 'EncryptionKeyId',
        'secretData'      => 'SecretData',
        'secretDataType'  => 'SecretDataType',
        'description'     => 'Description',
        'tags'            => 'Tags',
    ];

    public function validate()
    {
        Model::validateRequired('secretName', $this->secretName, true);
        Model::validateRequired('versionId', $this->versionId, true);
        Model::validateRequired('secretData', $this->secretData, true);
    }

    public function toMap()
    {
        $res = [];
        if (null !== $this->secretName) {
            $res['SecretName'] = $this->secretName;
        }
        if (null !== $this->versionId) {
            $res['VersionId'] = $this->versionId;
        }
        if (null !== $this->encryptionKeyId) {
            $res['EncryptionKeyId'] = $this->encryptionKeyId;
        }
        if (null !== $this->secretData) {
            $res['SecretData'] = $this->secretData;
        }
        if (null !== $this->secretDataType) {
            $res['SecretDataType'] = $this->secretDataType;
        }
        if (null !== $this->description) {
            $res['Description'] = $this->description;
        }
        if (null !== $this->tags) {
            $res['Tags'] = $this->tags;
        }

        return $res;
    }

    /**
     * @param array $map
     *
     * @return CreateSecretRequest
     */
    public static function fromMap($map = [])
    {
        $model = new self();
        if (isset($map['SecretName'])) {
            $model->secretName = $map['SecretName'];
        }
        if (isset($map['VersionId'])) {
            $model->versionId = $map['VersionId'];
        }
        if (isset($map['EncryptionKeyId'])) {
            $model->encryptionKeyId = $map['EncryptionKeyId'];
        }
        if (isset($map['SecretData'])) {
            $model->secretData = $map['SecretData'];
        }
        if (isset($map['SecretDataType'])) {
            $model->secretDataType = $map['SecretDataType'];
        }
        if (isset($map['Description'])) {
            $model->description = $map['Description'];
        }
        if (isset($map['Tags'])) {
            $model->tags = $map['Tags'];
        }

        return $model;
    }
}
