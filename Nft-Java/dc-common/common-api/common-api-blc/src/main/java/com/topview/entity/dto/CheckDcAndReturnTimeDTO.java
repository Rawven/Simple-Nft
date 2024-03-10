package com.topview.entity.dto;

import java.util.List;
import lombok.Data;
import lombok.experimental.Accessors;

/**
 * check dc and return time dto
 *
 * @author 刘家辉
 * @date 2024/03/09
 */
@Data
@Accessors(chain = true)
public class CheckDcAndReturnTimeDTO {
    private String owner;
    private List<byte[]> collectionHash;
}
