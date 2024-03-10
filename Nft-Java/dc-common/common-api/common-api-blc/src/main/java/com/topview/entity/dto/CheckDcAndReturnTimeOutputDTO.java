package com.topview.entity.dto;

import java.math.BigInteger;
import java.util.List;
import lombok.Data;
import lombok.experimental.Accessors;

/**
 * check dc and return time output dto
 *
 * @author 刘家辉
 * @date 2024/03/09
 */
@Data
@Accessors(chain = true)
public class CheckDcAndReturnTimeOutputDTO {
    private Boolean checkResult;
    private List<BigInteger> timeList;

    public CheckDcAndReturnTimeOutputDTO(Boolean value1, List<BigInteger> value2) {
        this.checkResult = value1;
        this.timeList = value2;
    }
}
