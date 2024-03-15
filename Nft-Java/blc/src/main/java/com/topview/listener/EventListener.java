package com.topview.listener;

import com.topview.client.ChainClient;
import com.topview.constant.BlcToUserTag;
import com.topview.constant.MqConstant;
import com.topview.entity.struct.DataStruct;
import com.topview.event.BlcNotice;
import com.topview.util.JsonUtil;
import com.topview.util.MqUtil;
import java.io.IOException;
import java.math.BigInteger;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.time.LocalDateTime;
import java.util.List;
import javax.annotation.PostConstruct;
import javax.annotation.Resource;
import lombok.extern.slf4j.Slf4j;
import org.fisco.bcos.sdk.jni.common.JniException;
import org.fisco.bcos.sdk.v3.client.Client;
import org.fisco.bcos.sdk.v3.codec.ContractCodec;
import org.fisco.bcos.sdk.v3.codec.ContractCodecException;
import org.fisco.bcos.sdk.v3.codec.abi.tools.TopicTools;
import org.fisco.bcos.sdk.v3.eventsub.EventSubParams;
import org.fisco.bcos.sdk.v3.eventsub.EventSubscribe;
import org.fisco.bcos.sdk.v3.model.EventLog;
import org.springframework.cloud.stream.function.StreamBridge;
import org.springframework.stereotype.Component;

import static com.topview.constant.EventConstant.CREATE_ACTIVITY;
import static com.topview.constant.EventConstant.CREATE_POOL;
import static com.topview.constant.EventConstant.POOL_TYPE;

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
    @Resource
    private StreamBridge streamBridge;
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
        poolLogicAbi = new String(Files.readAllBytes(Paths.get("Nft-Java/blc/src/main/resources/solidity/abi/PoolLogic.abi")));
        //开启监听
        //TODO 待优化
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
        // 设置合约事件
        params.addTopic(0, "LogCreatePool(uint256,address, DataStructs.Pool,uint256)");

        params.addTopic(0, "LogCreateActivity(uint256,address, DataStructs.Activity,uint256)");
        //初始化topic生成工具
        TopicTools tools = new TopicTools(client.getCryptoSuite());
        new TopicTools(client.getCryptoSuite());
        //计算默认topic
        String createPoolTopic = tools.stringToTopic(CREATE_POOL);
        String createActivityTopic = tools.stringToTopic(CREATE_ACTIVITY);
        //开始监听
        eventSubscribe.subscribeEvent(params, (eventSubId, status, logs) -> {
            for (EventLog event : logs) {
                String top = event.getTopics().get(0);
                try {
                    if (top.equals(createPoolTopic)) {
                        logCreatePool(event);
                    } else if (top.equals(createActivityTopic)) {
                        logCreateActivity(event);
                    } else {
                        log.error("未知的事件");
                    }
                } catch (ContractCodecException e) {
                    log.error("解析事件失败", e);
                    throw new RuntimeException("解析事件失败");
                }
            }
            System.out.println("event sub id: " + eventSubId);
            System.out.println(" \t status: " + status);
            System.out.println(" \t logs: " + logs);
        });
    }

    private void logCreatePool(EventLog event) throws ContractCodecException {
        String eventName = "LogCreatePool";
        /*
         * 第一个元素为对应的poolId
         * 第二个元素为操作者地址
         * 第三个元素为pool结构体
         */
        List<String> list = abiCodec.decodeEventToString(poolLogicAbi, eventName, event);
        //解析pool结构体字符串
        DataStruct.Pool structArgs = JsonUtil.jsonToObj(list.get(2), DataStruct.Pool.class);

        String description = "地址为" + list.get(1) + "的用户发行了【" + structArgs.getName() + "】的数藏池。" +
            "发行数量为" + structArgs.getAmount() + ",每位用户限制购买的数量为" + structArgs.getLimitAmount() + "。";
        BlcNotice notice = new BlcNotice()
            .setTitle("发行数藏池:" + structArgs.getName())
            .setType(POOL_TYPE)
            .setPublishTime(LocalDateTime.now())
            .setUserAddress(list.get(1))
            .setDescription(description);
        //发送通知消息
        streamBridge.send(MqConstant.CHANEL_FIRST, MqUtil.createMsg(JsonUtil.objToJson(notice), BlcToUserTag.BLC_NOTICE));
    }

    private void logCreateActivity(EventLog event) throws ContractCodecException {
        String eventName = "LogCreateActivity";
        List<String> list = abiCodec.decodeEventToString(poolLogicAbi, eventName, event);

        //解析activity结构体字符串
        DataStruct.Activity structArgs = JsonUtil.jsonToObj(list.get(2), DataStruct.Activity.class);

        String description = "地址为" + list.get(1) + "的用户举行了【" + structArgs.getName() + "】活动。" +
            "总数量为" + list.get(3) + "。";
        //发送通知消息
        BlcNotice notice = new BlcNotice()
            .setTitle("创建活动:" + structArgs.getName())
            .setType(POOL_TYPE)
            .setPublishTime(LocalDateTime.now())
            .setUserAddress(list.get(1))
            .setDescription(description);
        streamBridge.send(MqConstant.CHANEL_FIRST, MqUtil.createMsg(JsonUtil.objToJson(notice), BlcToUserTag.BLC_NOTICE));
    }
}
