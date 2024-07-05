CREATE TABLE `user`
(
    `id` bigint not null primary key auto_increment comment 'id',
    `username` varchar(30) not null comment '用户名',
    `password` varchar(255) not null comment '密码',
    `name` varchar(30) not null comment '名字',
    'lab_name' varchar(50) not null comment '实验室名称',
    `share_list` json comment '分享的paper',
    `paper_list` json comment '收到的paper'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;