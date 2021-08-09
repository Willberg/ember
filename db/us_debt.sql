create table `us_debt`
(
    `id`          bigint(20)     not null auto_increment comment 'id',
    `date`        varchar(20)    not null comment '日期',
    `debt`        decimal(50, 4) not null comment '债务',
    `create_time` bigint(20)     not null comment '创建时间',
    primary key (`id`),
    unique key (`date`)
) engine = innodb
  auto_increment = 1;