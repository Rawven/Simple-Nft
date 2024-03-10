package com.topview.entity.dto;

import java.math.BigInteger;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.experimental.Accessors;

/**
 * before mint dto
 *
 * @author 刘家辉
 * @date 2024/03/08
 */
@Data
@Accessors(chain = true)
@AllArgsConstructor
public class BeforeMintDTO {
    private BigInteger dcId;
    private byte[] uniqueId;

}
