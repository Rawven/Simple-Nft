package com.topview.event;

import lombok.Data;
import lombok.experimental.Accessors;

import java.math.BigInteger;

/**
 * createo pool notice
 *
 * @author 刘家辉
 * @date 2024/03/09
 */
@Data
@Accessors(chain = true)
public class BlcNotice {
    private String title;
    private Integer type;
    private BigInteger publishTime;
    private String userAddress;
    private String description;

}
