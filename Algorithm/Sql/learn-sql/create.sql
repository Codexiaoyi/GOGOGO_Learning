create table if not exists `runoob_tbl`(
    `runoob_id` int unsigned auto_increment,
    `runoob_title` varchar(100) not null,
    `runoob_author` varchar(40) not null,
    `submission_date` Date,
    primary key (`runoob_id`)
)engine=InnoDB default charset=utf8;