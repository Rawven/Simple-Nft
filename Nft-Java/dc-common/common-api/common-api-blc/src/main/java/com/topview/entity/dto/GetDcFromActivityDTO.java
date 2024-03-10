package com.topview.entity.dto;

import java.math.BigInteger;
import lombok.Data;
import lombok.experimental.Accessors;

/**
 * get dc from activity dto
 *
 * @author 刘家辉
 * @date 2024/03/08
 */
@Data
@Accessors(chain = true)
public class GetDcFromActivityDTO {
    private BigInteger activityId;
    private byte[] password;

}
