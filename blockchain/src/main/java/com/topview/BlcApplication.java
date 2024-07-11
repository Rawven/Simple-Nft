package com.topview;

import org.apache.dubbo.config.spring.context.annotation.EnableDubbo;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.EnableAspectJAutoProxy;

/**
 * blc application
 *
 * @author 刘家辉
 * @date 2024/03/09
 */
@SpringBootApplication
@EnableAspectJAutoProxy
@EnableDubbo(scanBasePackages = "com.topview.api")
public class BlcApplication {
    public static void main(String[] args) {
        SpringApplication.run(BlcApplication.class, args);
    }
}
