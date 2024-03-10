package com.topview.entity.dto;

import java.math.BigInteger;
import lombok.Data;
import lombok.experimental.Accessors;

/**
 * create activity dto
 *
 * @author 刘家辉
 * @date 2024/03/08
 */
@Data
@Accessors(chain = true)
public class CreateActivityDTO {
    private String name;
    private String password;
    private BigInteger amount;
    private String cid;
    private String dcName;
}
