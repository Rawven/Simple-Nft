package com.topview.util;

import cn.hutool.core.util.IdUtil;
import com.topview.constant.MqConstant;
import com.topview.event.Event;
import java.util.HashMap;
import java.util.Map;
import java.util.concurrent.TimeUnit;
import lombok.extern.slf4j.Slf4j;
import org.apache.rocketmq.common.message.MessageConst;
import org.springframework.messaging.Message;
import org.springframework.messaging.support.GenericMessage;

import static com.topview.constant.MqConstant.HEAD;

/**
 * mq util
 *
 * @author 刘家辉
 * @date 2023/12/08
 */
@Slf4j
public class MqUtil {

    public static Message<Event> createMsg(String data, String tag) {
        Map<String, Object> headers = new HashMap<>(2);
        headers.put(MessageConst.PROPERTY_KEYS, IdUtil.getSnowflakeNextIdStr());
        headers.put(MessageConst.PROPERTY_TAGS, tag);
        return new GenericMessage<>(
            new Event(data), headers);
    }





}
