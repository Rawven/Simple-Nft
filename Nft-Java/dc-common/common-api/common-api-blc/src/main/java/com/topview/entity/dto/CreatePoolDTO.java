package com.topview.entity.dto;

import lombok.Data;
import lombok.experimental.Accessors;

/**
 * create pool dto
 *
 * @author 刘家辉
 * @date 2024/03/09
 */
@Data
@Accessors(chain = true)
public class CreatePoolDTO {
    private Integer limitAmount;
    private Integer price;
    private Integer amount;
    private String cid;
    private String dcName;
}
