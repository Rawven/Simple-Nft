package com.topview.event;

import java.time.LocalDateTime;
import lombok.Data;
import lombok.experimental.Accessors;

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
    private LocalDateTime publishTime;
    private String userAddress;
    private String description;

}
