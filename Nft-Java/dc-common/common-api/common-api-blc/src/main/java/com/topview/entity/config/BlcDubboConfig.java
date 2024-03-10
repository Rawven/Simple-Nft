package com.topview.entity.config;

import com.topview.entity.api.BlcRpcService;
import org.apache.dubbo.config.annotation.DubboReference;
import org.apache.dubbo.config.spring.ReferenceBean;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

/**
 * dubbo config
 *
 * @author 刘家辉
 * @date 2024/02/06
 */
@Configuration
public class BlcDubboConfig {
    @Bean
    @DubboReference(interfaceClass = BlcRpcService.class, version = "1.0.0", timeout = 15000)
    public ReferenceBean<BlcRpcService> userRpcService() {
        return new ReferenceBean<>();
    }
}
