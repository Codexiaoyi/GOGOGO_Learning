-- UNION 操作符用于连接两个以上的 SELECT 语句的结果组合到一个结果集合中。多个 SELECT 语句会删除重复的数据。
-- 前置预备表
create table if not exists `Websites`(
    `id` int unsigned,
    `name` varchar(40) not null,
    `url` varchar(100) not null,
    `alexa` int,
    `country` varchar(20) not null,
    primary key (`id`)
) engine = InnoDB default charset = utf8;

create table if not exists `apps`(
    `id` int unsigned,
    `app_name` varchar(40) not null,
    `url` varchar(100) not null,
    `country` varchar(20) not null,
    primary key (`id`)
) engine = InnoDB default charset = utf8;

-- 插入数据
insert into
    Websites
values
    (1, 'Google', 'google', 1, 'USA'),
    (2, '淘宝', 'taobao', 13, 'CN'),
    (3, '菜鸟教程', 'runoob', 4689, 'CN'),
    (4, 'Facebook', 'facebook', 3, 'USA');

insert into
    apps
values
    (1, 'QQ', 'qq', 'CN'),
    (2, '微博', 'weibo', 'CN'),
    (3, '淘宝', 'taobao', 'CN');

-- union
select
    country
from
    Websites
union
select
    country
from
    apps
order by
    country;

-- union2
select
    country,
    name as combine_name
from
    Websites
where
    country = 'CN'
union
all
select
    country,
    app_name
from
    apps
where
    country = 'CN'
order by
    country;

-- drop
drop table Websites;

drop table apps;