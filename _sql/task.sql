USE scheduler;
CREATE TABLE task (
  `id` VARCHAR(128) NOT NULL,
  `name` VARCHAR(128) NOT NULL,
  `task_type` VARCHAR(128) NOT NULL,
  `input` TEXT NOT NULL,
  `output` TEXT NOT NULL,
  `commit_time` BIGINT,
  `begin_time` BIGINT,
  `end_time` BIGINT,
  `status` VARCHAR(32),
  `worker_id` VARCHAR(128),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;