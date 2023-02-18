use scheduler;
create table if not exists task (
  `id` varchar(128) not null Comment 'task id, the primary key of a task',
  `name` varchar(128) not null Comment 'task name', 
  `task_type` varchar(128) not null Comment 'task type',
  `input` text Comment 'task input',
  `output` text Comment 'task output',
  `commit_time` bigint(20) Comment 'the time (ms) of the user apply the task',
  `begin_time` bigint(20) Comment 'the time (ms) of the user begin the task',
  `end_time` bigint(20) Comment 'the time (ms) of the user finished the task',
  `status` varchar(32) Comment 'the task status',
  `worker_id` bigint(20) Comment 'the worker_id which execute this task',
  primary key (`id`)
) CHARACTER SET utf8 COLLATE utf8_general_ci;