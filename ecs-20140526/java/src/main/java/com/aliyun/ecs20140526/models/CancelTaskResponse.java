// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.ecs20140526.models;

import com.aliyun.tea.*;

public class CancelTaskResponse extends TeaModel {
    @NameInMap("RequestId")
    @Validation(required = true)
    public String requestId;

    public static CancelTaskResponse build(java.util.Map<String, ?> map) throws Exception {
        CancelTaskResponse self = new CancelTaskResponse();
        return TeaModel.build(map, self);
    }

}
