package com.topview.client;

import cn.hutool.core.convert.Convert;
import com.github.benmanes.caffeine.cache.Cache;
import com.github.benmanes.caffeine.cache.Caffeine;
import com.topview.blc.PoolData;
import com.topview.blc.PoolLogic;
import com.topview.blc.UserData;
import com.topview.blc.UserLogic;
import com.topview.config.ContractProperty;
import java.io.InputStream;
import javax.annotation.PostConstruct;
import lombok.Data;
import lombok.extern.slf4j.Slf4j;
import org.fisco.bcos.sdk.v3.BcosSDK;
import org.fisco.bcos.sdk.v3.client.Client;
import org.fisco.bcos.sdk.v3.config.ConfigOption;
import org.fisco.bcos.sdk.v3.config.exceptions.ConfigException;
import org.fisco.bcos.sdk.v3.config.model.ConfigProperty;
import org.fisco.bcos.sdk.v3.crypto.CryptoSuite;
import org.fisco.bcos.sdk.v3.crypto.keypair.CryptoKeyPair;
import org.fisco.bcos.sdk.v3.model.CryptoType;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;
import org.yaml.snakeyaml.Yaml;
import org.yaml.snakeyaml.representer.Representer;

/**
 * FiscoBcos配置类
 *
 * @author ashinnotfound
 * @date 2023/01/25
 */
@Data
@Component
@Slf4j
public class ChainClient {
    private BcosSDK sdk;
    private Client client;
    private CryptoKeyPair adminKeyPair;
    private CryptoSuite cryptoSuite;
    private Cache<Object, Object> localCache;
    @Value("${fisco.file-path}")
    private String fiscoConfigPath;
    @Value("${fisco.group}")
    private String groupId;
    @Value("${fisco.contract.admin}")
    private String adminPrivateKey;
    @Autowired
    private ContractProperty contractProperty;

    @PostConstruct
    public void init() {
        try {
            //初始化sdk
            Representer representer = new Representer();
            representer.getPropertyUtils().setSkipMissingProperties(true);
            Yaml yaml = new Yaml(representer);
            InputStream inputStream = this.getClass().getResourceAsStream(fiscoConfigPath);
            ConfigProperty configProperty = yaml.loadAs(inputStream, ConfigProperty.class);
            ConfigOption configOption = new ConfigOption(configProperty);
            sdk = new BcosSDK(configOption);
            cryptoSuite = new CryptoSuite(CryptoType.ECDSA_TYPE);
            // 从十六进制私钥字符串hexPrivateKey加载私钥对象
            adminKeyPair = cryptoSuite.getKeyPairFactory().createKeyPair(adminPrivateKey);
            client = sdk.getClient(groupId);
            localCache = Caffeine.newBuilder()
                .maximumSize(500)
                .initialCapacity(10)
                .recordStats().build();
        } catch (ConfigException e) {
            log.error("sdk初始化失败", e);
        }
    }

    public <T> T getContractAdminInstance(Class<T> tClass) {
        return getContractInstance(tClass, adminKeyPair.getHexPrivateKey());
    }

    public <T> T getContractInstance(Class<T> tClass, String userKey) {
        userKey = userKey.replace("\\", "").replace("\"", "");
        log.info("userKey:{}", userKey);
        Object contract = localCache.getIfPresent(userKey + "_" + tClass.getName());
        if (contract != null) {
            return Convert.convert(tClass, contract);
        } else {
            if (tClass == PoolData.class) {
                CryptoKeyPair cryptoKeyPair = cryptoSuite.getKeyPairFactory().createKeyPair(userKey);
                PoolData load = PoolData.load(contractProperty.poolDataAddress, client, cryptoKeyPair);
                localCache.put(userKey + "_" + tClass.getName(), load);
                return Convert.convert(tClass, load);
            } else if (tClass == (PoolLogic.class)) {
                CryptoKeyPair cryptoKeyPair = cryptoSuite.getKeyPairFactory().createKeyPair(userKey);
                PoolLogic load = PoolLogic.load(contractProperty.poolLogicAddress, client, cryptoKeyPair);
                localCache.put(userKey + "_" + tClass.getName(), load);
                return Convert.convert(tClass, load);
            } else if (tClass == (UserLogic.class)) {
                CryptoKeyPair cryptoKeyPair = cryptoSuite.getKeyPairFactory().createKeyPair(userKey);
                UserLogic load = UserLogic.load(contractProperty.userLogicAddress, client, cryptoKeyPair);
                localCache.put(userKey + "_" + tClass.getName(), load);
                return Convert.convert(tClass, load);
            } else if (tClass == (UserData.class)) {
                CryptoKeyPair cryptoKeyPair = cryptoSuite.getKeyPairFactory().createKeyPair(userKey);
                UserData load = UserData.load(contractProperty.userDataAddress, client, cryptoKeyPair);
                localCache.put(userKey + "_" + tClass.getName(), load);
                return Convert.convert(tClass, load);
            } else {
                throw new RuntimeException("未知合约");
            }
        }
    }
}
