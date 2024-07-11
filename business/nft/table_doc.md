CREATE TABLE pool (
pool_id INT AUTO_INCREMENT PRIMARY KEY COMMENT '池id',

                      cid VARCHAR(255) COMMENT 'ipfs上图片存储地址',

                      name VARCHAR(255) COMMENT '名字',

                      description VARCHAR(255) COMMENT '描述',

                      price INT COMMENT '价格',

                      amount INT COMMENT '数量',

                      `left` INT COMMENT '剩余数量',

                      limit_amount INT COMMENT '限制数量',

                      creator_name VARCHAR(255) COMMENT '创建者名字',

                      creator_address VARCHAR(255) COMMENT '创建者地址',

                      status BOOLEAN COMMENT '状态（true=>新品，false=>非卖）'

) COMMENT='池信息表';

CREATE TABLE dc (
id INT AUTO_INCREMENT PRIMARY KEY COMMENT '唯一标识ID',

                    hash VARCHAR(255) COMMENT '藏品hash',

                    cid VARCHAR(255) COMMENT 'ipfs上图片存储地址',

                    name VARCHAR(255) COMMENT '名字',

                    description VARCHAR(255) COMMENT '描述',

                    price INT COMMENT '价格',

                    owner_name VARCHAR(255) COMMENT '拥有者名字',

                    owner_address VARCHAR(255) COMMENT '拥有者地址',

                    creator_name VARCHAR(255) COMMENT '创建者名字',

                    creator_address VARCHAR(255) COMMENT '创建者地址'

) COMMENT='藏品信息表';

CREATE TABLE activity (
id INT AUTO_INCREMENT PRIMARY KEY COMMENT '链中返回的id',

                          name VARCHAR(255) COMMENT '活动名',

                          description VARCHAR(255) COMMENT '活动描述',

                          dc_description VARCHAR(255) COMMENT 'nft描述',

                          cid VARCHAR(255) COMMENT '图片存储地址',

                          host_name VARCHAR(255) COMMENT '举办人名字',

                          host_address VARCHAR(255) COMMENT '举办人地址',

                          amount INT COMMENT '总数量',

                          remainder INT COMMENT '剩余数量',

                          status BOOLEAN COMMENT '状态(0=>展示，1=>不展示)'

) COMMENT='活动信息表';
