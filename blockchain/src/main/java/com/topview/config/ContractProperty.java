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
    @Value("${fisco.contract.address.poolData}")
    public String poolDataAddress;
    @Value("${fisco.contract.address.poolLogic}")
    public String poolLogicAddress;
    @Value("${fisco.contract.address.userData}")
    public String userDataAddress;
    @Value("${fisco.contract.address.userLogic}")
    public String userLogicAddress;

}
