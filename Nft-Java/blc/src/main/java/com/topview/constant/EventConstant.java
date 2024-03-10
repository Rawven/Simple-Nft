package com.topview.constant;

/**
 * 事件签名常量
 *
 * @author lql
 * @date 2023/04/20
 */
public class EventConstant {
    public final static String CREATE_POOL = "LogCreatePool(uint256,address,(string,string,uint256,uint256,uint256,uint256,address,uint256))";

    public final static String CREATE_ACTIVITY = "LogCreateActivity(uint256,address,(string,bytes,uint256),uint256)";

    public final static Integer POOL_TYPE = 1;

    public final static Integer ACTIVITY_TYPE = 0;
}
