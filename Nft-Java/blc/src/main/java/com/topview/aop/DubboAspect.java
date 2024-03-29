package com.topview.aop;

import java.util.Arrays;
import lombok.extern.slf4j.Slf4j;
import org.aspectj.lang.ProceedingJoinPoint;
import org.aspectj.lang.annotation.Around;
import org.aspectj.lang.annotation.Aspect;
import org.aspectj.lang.annotation.Pointcut;
import org.springframework.stereotype.Component;

/**
 * dubbo aspect
 *
 * @author 刘家辉
 * @date 2024/01/20
 */
@Aspect
@Component
@Slf4j
public class DubboAspect {
    @Pointcut("execution(* com.topview.api..*.*(..))")
    public void pointcut() {
    }

    @Around("pointcut()")
    public Object around(ProceedingJoinPoint joinPoint) throws Throwable {
        // 获取方法的全名，包括类名和方法名
        String methodFullName = joinPoint.getSignature().toLongString();
        // 记录方法开始执行的时间
        long startTime = System.currentTimeMillis();
        // 执行方法
        Object result = joinPoint.proceed();
        // 记录方法结束执行的时间
        long endTime = System.currentTimeMillis();
        // 计算方法执行的时间
        long executeTime = endTime - startTime;
        // 打印出方法的全名、结果和执行时间
        log.info("--Blc服务被调用 : 方法全名: {}, 结果: {}, 执行时间: {} 毫秒", methodFullName, result, executeTime);
        return result;

    }

}
