package com.topview.entity.dto;

import lombok.Data;
import lombok.experimental.Accessors;

/**
 * user key
 *
 * @author 刘家辉
 * @date 2024/03/08
 */
@Data
@Accessors(chain = true)
public class UserKey {
    /**
     * user key
     */
    private String userKey;
}
