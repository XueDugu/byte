create table douyin.comments
(
    id         bigint auto_increment comment '自增主键'
        primary key,
    user_id    bigint                            not null comment '用户id',
    video_id   bigint                            not null comment '视频id',
    text       varchar(4096) collate utf8mb4_bin null comment '评论文本内容',
    created_at datetime                          null comment '创建时间',
    deleted_at datetime                          null comment '删除时间'
)
    comment '评论表' collate = utf8_bin;

create table douyin.favorites
(
    id       bigint auto_increment comment '自增主键'
        primary key,
    user_id  bigint  not null comment '用户id',
    video_id bigint  not null comment '视频id',
    status   tinyint not null comment '点赞状态（1点赞，2取消点赞）',
    constraint favorites_pk2
        unique (user_id, video_id)
)
    comment '点赞表';

create table douyin.messages
(
    id         bigint auto_increment comment '自增主键'
        primary key,
    from_id    bigint                not null comment '消息发送方的id',
    to_id      bigint                not null comment '消息接收方的id',
    content    text collate utf8_bin null comment '消息内容',
    created_at datetime              null comment '创建时间'
)
    comment '用户之间发送的消息';

create table douyin.relations
(
    id      bigint auto_increment comment '自增主键'
        primary key,
    fan_id  bigint not null comment '粉丝的用户id',
    user_id bigint not null comment '被关注用户的id',
    status  int    not null comment '关注状态（1关注，2取关）',
    constraint relations_pk
        unique (user_id, fan_id)
)
    comment '用户之间的关注关系';

create table douyin.users
(
    id         bigint auto_increment comment '自增主键'
        primary key,
    name       varchar(8) not null comment '用户名',
    password   varchar(8) null comment '密码',
    created_at datetime   null,
    updated_at datetime   null,
    constraint users_name
        unique (name)
)
    comment '用户表';

create table douyin.videos
(
    id         bigint auto_increment comment '自增主键'
        primary key,
    title      varchar(16)  null comment '视频名称',
    author     bigint       null comment '作者用户id（外键）',
    play_url   varchar(128) not null,
    cover_url  varchar(128) null,
    created_at datetime     null
);

