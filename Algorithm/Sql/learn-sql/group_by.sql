SET
    NAMES utf8;

SET
    FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `employee_tbl`
-- ----------------------------
DROP TABLE IF EXISTS `employee_tbl`;

CREATE TABLE `employee_tbl` (
    `id` int(11) NOT NULL,
    `name` char(10) NOT NULL DEFAULT '',
    `date` datetime NOT NULL,
    `signin` tinyint(4) NOT NULL DEFAULT '0' COMMENT '登录次数',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

-- ----------------------------
--  Records of `employee_tbl`
-- ----------------------------
BEGIN;

INSERT INTO
    `employee_tbl`
VALUES
    ('1', '小明', '2016-04-22 15:25:33', '1'),
    ('2', '小王', '2016-04-20 15:25:47', '3'),
    ('3', '小丽', '2016-04-19 15:26:02', '2'),
    ('4', '小王', '2016-04-07 15:26:14', '4'),
    ('5', '小明', '2016-04-11 15:26:40', '4'),
    ('6', '小明', '2016-04-04 15:26:54', '2');

COMMIT;

SET
    FOREIGN_KEY_CHECKS = 1;

-- 按照名字分组，统计每个人有几条记录
select
    name,
    count(*) as count
from
    employee_tbl
group by
    name;

-- with rollup可以在分组统计基础上再进行相同字段统计
select
    name,
    sum(signin) as signin_count
from
    employee_tbl
group by
    name with rollup;

-- coalesce(a,b,c) a==null，选择b; b==null，选择c...
select
    coalesce(name, '总数') as 'name/总数',
    sum(signin) as signin_count
from
    employee_tbl
group by
    name with rollup;