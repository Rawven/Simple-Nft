package com.topview.util;

import cn.hutool.core.util.IdUtil;
import lombok.extern.slf4j.Slf4j;
import org.apache.rocketmq.common.message.MessageConst;
import org.springframework.messaging.Message;
import org.springframework.messaging.support.GenericMessage;

import java.util.HashMap;
import java.util.Map;

/**
 * mq util
 *
 * @author 刘家辉
 * @date 2023/12/08
 */
@Slf4j
public class MqUtil {

    public static Message<String> createMsg(String data) {
        Map<String, Object> headers = new HashMap<>(1);
        headers.put(MessageConst.PROPERTY_KEYS, IdUtil.getSnowflakeNextIdStr());
        return new GenericMessage<>(
                data, headers);
    }

}
