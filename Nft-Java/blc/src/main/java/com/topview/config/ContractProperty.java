package com.topview.config;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Configuration;

/**
 * contract property
 *
 * @author 刘家辉
 * @date 2024/03/08
 */
@Configuration

public class ContractProperty {
    @Value("${contract.address.poolData}")
    public static String poolDataAddress;
    @Value("${contract.address.poolLogic}")
    public static String poolLogicAddress;
    @Value("${contract.address.userData}")
    public static String userDataAddress;
    @Value("${contract.address.userLogic}")
    public static String userLogicAddress;

}
