package com.topview.util;

import cn.hutool.core.util.IdUtil;
import com.topview.constant.MqConstant;
import com.topview.event.Event;
import java.util.HashMap;
import java.util.Map;
import java.util.concurrent.TimeUnit;
import lombok.extern.slf4j.Slf4j;
import org.apache.rocketmq.common.message.MessageConst;
import org.redisson.api.RedissonClient;
import org.springframework.data.redis.core.StringRedisTemplate;
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

    public static boolean checkMsgIsvalid(Message<Event> msg, RedissonClient redissonClient) {
        Object id = msg.getHeaders().get(MqConstant.HEADER_KEYS);
        if (id == null || redissonClient.getBucket(HEAD + id).isExists()) {
            log.info("--RocketMq 重复或非法的消息，不处理");
            return true;
        }
        return false;
    }

    public static void protectMsg(Message<Event> msg, RedissonClient redissonClient) {
        Object id = msg.getHeaders().get(MqConstant.HEADER_KEYS);
        redissonClient.getBucket(HEAD + id).set(id, MqConstant.EXPIRE_TIME, TimeUnit.MINUTES);
    }

    public static boolean checkMsgIsvalid(Message<Event> msg, StringRedisTemplate template) {
        Object id = msg.getHeaders().get(MqConstant.HEADER_KEYS);
        if (id == null || Boolean.TRUE.equals(template.hasKey(HEAD + id))) {
            log.info("--RocketMq 重复或非法的消息，不处理");
            return true;
        }
        return false;
    }

    public static void protectMsg(Message<Event> msg, StringRedisTemplate template) {
        Object id = msg.getHeaders().get(MqConstant.HEADER_KEYS);
        if (id == null) {
            throw new RuntimeException("消息id为空");
        }
        template.opsForValue().set(HEAD + id, id.toString(), MqConstant.EXPIRE_TIME, TimeUnit.MINUTES);
    }
}
