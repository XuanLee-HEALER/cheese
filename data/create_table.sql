-- cheese.role_url definition

CREATE TABLE `role_url` (
  `roleName` varchar(20) NOT NULL,
  `roleUrl` varchar(100) NOT NULL,
  PRIMARY KEY (`roleName`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色-详情url对照关系';