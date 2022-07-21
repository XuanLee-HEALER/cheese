-- cheese.role_url definition

CREATE TABLE `role_url` (
  `roleName` varchar(20) NOT NULL,
  `roleUrl` varchar(100) NOT NULL,
  PRIMARY KEY (`roleName`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色-详情url对照关系';

-- cheese.role definition
CREATE TABLE cheese.`role` (
	id BIGINT auto_increment NOT NULL,
	name varchar(10) NOT NULL COMMENT '名字',
	elementType SMALLINT NULL COMMENT '元素类型，枚举',
	birth varchar(100) NULL COMMENT '生日',
	fromWhere varchar(50) NULL COMMENT '归属地',
	weapon SMALLINT NOT NULL COMMENT '武器类型，枚举',
	destiny varchar(100) NULL COMMENT '命之座',
	dub varchar(20) NULL,
	CONSTRAINT role_PK PRIMARY KEY (id)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_general_ci;
