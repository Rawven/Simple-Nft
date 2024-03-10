package com.topview.entity.dto;

import com.topview.entity.struct.DataStruct;
import lombok.Data;
import lombok.experimental.Accessors;

@Data
@Accessors(chain = true)
public class ActivityAndPool {
    private DataStruct.Activity activity;
    private DataStruct.Pool pool;

    public ActivityAndPool(DataStruct.Activity value1, DataStruct.Pool value2) {
        this.activity = value1;
        this.pool = value2;
    }
}
