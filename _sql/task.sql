USE scheduler;
CREATE TABLE task (
  `id` VARCHAR(128) NOT NULL COMMENT 'task id, the primary key of a task',
  `name` VARCHAR(128) NOT NULL COMMENT 'task name',
  `task_type` VARCHAR(128) NOT NULL COMMENT 'the type of task',
  `input` TEXT COMMENT 'task input',
  `output` TEXT COMMENT 'task output',
  `commit_time` BIGINT COMMENT 'the time (ms) of the user apply the task',
  `begin_time` BIGINT COMMENT 'the time (ms) of the task begin',
  `end_time` BIGINT COMMENT 'the time (ms) of the task finished',
  `status` VARCHAR(32) COMMENT 'the task status',
  `worker_id` VARCHAR(128) COMMENT 'the worker_id which execute this task',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;