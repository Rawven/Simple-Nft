package com.topview.listener;

import cn.hutool.core.convert.Convert;
import com.topview.client.ChainClient;
import com.topview.constant.BlcToUserTag;
import com.topview.entity.struct.DataStruct;
import com.topview.event.BlcNotice;
import com.topview.util.JsonUtil;
import com.topview.util.MqUtil;
import lombok.extern.slf4j.Slf4j;
import org.apache.rocketmq.client.producer.SendCallback;
import org.apache.rocketmq.client.producer.SendResult;
import org.apache.rocketmq.spring.core.RocketMQTemplate;
import org.fisco.bcos.sdk.jni.common.JniException;
import org.fisco.bcos.sdk.v3.client.Client;
import org.fisco.bcos.sdk.v3.codec.ContractCodec;
import org.fisco.bcos.sdk.v3.codec.ContractCodecException;
import org.fisco.bcos.sdk.v3.codec.abi.tools.TopicTools;
import org.fisco.bcos.sdk.v3.eventsub.EventSubParams;
import org.fisco.bcos.sdk.v3.eventsub.EventSubscribe;
import org.fisco.bcos.sdk.v3.model.EventLog;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.messaging.Message;
import org.springframework.stereotype.Component;

import javax.annotation.PostConstruct;
import javax.annotation.Resource;
import java.io.IOException;
import java.math.BigInteger;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.List;

import static com.topview.constant.EventConstant.*;

/**
 * 公告服务impl
 *
 * @author lql
 * @date 2023/04/16
 */
@Slf4j
@Component
public class EventListener {
    @Resource
    private ChainClient sdk;
    @Autowired
    private RocketMQTemplate rocketmqTemplate;
    private Client client;

    private ContractCodec abiCodec;

    private String poolLogicAbi;

    @PostConstruct
    public void init() throws IOException {
        //获取群组客户端
        client = sdk.getClient();
        //加载abi解码工具
        abiCodec = new ContractCodec(client.getCryptoSuite(), false);
        //加载abi字符串
        poolLogicAbi = new String(Files.readAllBytes(Paths.get("blc/src/main/resources/solidity/abi/PoolLogic.abi")));
        //开启监听
        createListen();
    }

    private void createListen() {
        // 初始化EventSubscribe
        EventSubscribe eventSubscribe;
        try {
            eventSubscribe = EventSubscribe.build(sdk.getClient());
        } catch (JniException e) {
            log.error("初始化EventSubscribe失败", e);
            throw new RuntimeException("初始化EventSubscribe失败");
        }
        eventSubscribe.start();

        // 设置参数
        EventSubParams params = new EventSubParams();

        // 从订阅时的最新区块区块开始，fromBlock设置为-1
        params.setFromBlock(BigInteger.valueOf(-1));
        // toBlock设置为-1，处理至最新区块继续等待新的区块
        params.setToBlock(BigInteger.valueOf(-1));
        //初始化topic生成工具
        TopicTools tools = new TopicTools(client.getCryptoSuite());
        //计算默认topic
        String createPoolTopic = tools.stringToTopic(CREATE_POOL);
        String createActivityTopic = tools.stringToTopic(CREATE_ACTIVITY);
        // 设置合约事件
        params.addTopic(0, createPoolTopic);
        params.addTopic(0, createActivityTopic);
        //开始监听
        eventSubscribe.subscribeEvent(params, (eventSubId, status, logs) -> {
            for (EventLog event : logs) {
                String top = event.getTopics().get(0);
                try {
                    if (top.equals(createPoolTopic)) {
                        log.info("EventListener--创建数藏池事件");
                        logCreatePool(event);
                    } else if (top.equals(createActivityTopic)) {
                        log.info("EventListener--创建活动事件");
                        logCreateActivity(event);
                    } else {
                        log.error("未知的事件");
                    }
                } catch (ContractCodecException e) {
                    log.error("解析事件失败", e);
                    throw new RuntimeException("解析事件失败");
                }
            }
            log.info("event sub id: " + eventSubId);
            log.info(" \t status: " + status);
            log.info(" \t logs: " + logs);
        });
    }

    private void logCreatePool(EventLog event) throws ContractCodecException {
        String eventName = "LogCreatePool";
        List<String> list = abiCodec.decodeEventToString(poolLogicAbi, eventName, event);
        List<Object> pool = JsonUtil.strToList(list.get(2));
        DataStruct.Pool structArgs = new DataStruct.Pool()
                .setCid(String.valueOf(pool.get(0)))
                .setName(String.valueOf(pool.get(1)))
                .setPrice(new BigInteger(String.valueOf(pool.get(2))))
                .setAmount(new BigInteger(String.valueOf(pool.get(3))))
                .setLeft(new BigInteger(String.valueOf(pool.get(4))))
                .setLimitAmount(new BigInteger(String.valueOf(pool.get(5))))
                .setCreator(String.valueOf(pool.get(6)))
                .setCreateTime(new BigInteger(String.valueOf(pool.get(7))));
        String description = "地址为" + list.get(1) + "的用户发行了【" + structArgs.getName() + "】的数藏池。" +
            "发行数量为" + structArgs.getAmount() + ",每位用户限制购买的数量为" + structArgs.getLimitAmount() + "。";
        BlcNotice notice = new BlcNotice()
            .setTitle("发行数藏池:" + structArgs.getName())
            .setType(POOL_TYPE)
                .setPublishTime(BigInteger.valueOf(System.currentTimeMillis()))
            .setUserAddress(list.get(1))
            .setDescription(description);
        //发送通知消息
        sendBlcNotice(notice);
    }

    private void logCreateActivity(EventLog event) throws ContractCodecException {
        String eventName = "LogCreateActivity";
        List<String> list = abiCodec.decodeEventToString(poolLogicAbi, eventName, event);
        //解析activity结构体字符串
        List<Object> activity = JsonUtil.strToList(list.get(2));
        DataStruct.Activity structArgs = new DataStruct.Activity()
                .setName(String.valueOf(activity.get(0)))
                .setEncodedKey(Convert.toPrimitiveByteArray(activity.get(1)))
                .setPoolId(Convert.toBigInteger(activity.get(2)));
        String description = "地址为" + list.get(1) + "的用户举行了【" + structArgs.getName() + "】活动。" +
            "总数量为" + list.get(3) + "。";
        //发送通知消息
        BlcNotice notice = new BlcNotice()
            .setTitle("创建活动:" + structArgs.getName())
            .setType(POOL_TYPE)
                .setPublishTime(BigInteger.valueOf(System.currentTimeMillis()))
            .setUserAddress(list.get(1))
            .setDescription(description);
        //发送通知消息
        sendBlcNotice(notice);
    }

    private void sendBlcNotice(BlcNotice notice) {
        Message<String> msg = MqUtil.createMsg(JsonUtil.objToJson(notice));
        rocketmqTemplate.asyncSend("Nft-Go:" + BlcToUserTag.BLC_NOTICE, msg, new SendCallback() {
            @Override
            public void onSuccess(SendResult sendResult) {
                log.info("发送消息成功");
            }

            @Override
            public void onException(Throwable throwable) {
                log.error("发送消息失败", throwable);
            }
        });
    }

}
