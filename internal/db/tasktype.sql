use scheduler;
create table if not exists task_type (
  task_type_id varchar(64) not null comment 'task type id',
  task_type_name varchar(64) unique not null comment 'task type name',
  max_task_in_que_limit int(10) not null comment 'the max task waiting for scheduled, -1 for unlimited',
  create_time bigint(20) not null comment 'this task type create time(unix second)',
  extra_info blob not null,
  primary key (task_type_id)
) CHARACTER SET utf8 COLLATE utf8_general_ci;