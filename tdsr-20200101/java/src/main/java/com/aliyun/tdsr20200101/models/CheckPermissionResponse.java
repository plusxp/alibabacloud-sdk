// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.tdsr20200101.models;

import com.aliyun.tea.*;

public class CheckPermissionResponse extends TeaModel {
    @NameInMap("RequestId")
    @Validation(required = true)
    public String requestId;

    @NameInMap("Success")
    @Validation(required = true)
    public Boolean success;

    @NameInMap("ErrMessage")
    @Validation(required = true)
    public String errMessage;

    public static CheckPermissionResponse build(java.util.Map<String, ?> map) throws Exception {
        CheckPermissionResponse self = new CheckPermissionResponse();
        return TeaModel.build(map, self);
    }

}
