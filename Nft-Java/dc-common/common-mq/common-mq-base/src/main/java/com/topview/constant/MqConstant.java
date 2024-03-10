package com.topview.constant;

/**
 * mq constant
 *
 * @author 刘家辉
 * @date 2023/12/10
 */
public class MqConstant {
    /**
     * 过期时间默认10分钟
     */
    public final static long EXPIRE_TIME = 10;

    public static final String HEAD = "rocketMq_msg_";
    public static final String HEADER_KEYS = "ROCKET_KEYS";
    public static final String HEADER_TAGS = "ROCKET_TAGS";

    public static final String CHANEL_FIRST = "producer-out-0";
    public static final String CHANEL_SECOND = "producer-out-1";

    public static final String SEARCH_POOL_KEY = "searchPool";
}
