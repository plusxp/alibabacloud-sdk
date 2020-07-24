// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.pts20190810.models;

import com.aliyun.tea.*;

public class SaveSceneResponse extends TeaModel {
    @NameInMap("RequestId")
    @Validation(required = true)
    public String requestId;

    @NameInMap("SceneId")
    @Validation(required = true)
    public String sceneId;

    @NameInMap("Message")
    @Validation(required = true)
    public String message;

    @NameInMap("Code")
    @Validation(required = true)
    public String code;

    @NameInMap("Success")
    @Validation(required = true)
    public Boolean success;

    @NameInMap("HttpStatusCode")
    @Validation(required = true)
    public Integer httpStatusCode;

    @NameInMap("CodeKey")
    @Validation(required = true)
    public String codeKey;

    public static SaveSceneResponse build(java.util.Map<String, ?> map) throws Exception {
        SaveSceneResponse self = new SaveSceneResponse();
        return TeaModel.build(map, self);
    }

}
