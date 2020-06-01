CREATE TABLE `t_export_record` (
	`f_id` VARCHAR(255) PRIMARY KEY,
	`f_staff_name` VARCHAR(255) NOT NULL,
	`f_card_id` VARCHAR(255) NOT NULL,
	`f_batch_flag` VARCHAR(255) NOT NULL,
	`f_create_time` TIMESTAMP NOT NULL,
	`f_update_time` TIMESTAMP NOT NULL
) ENGINE=INNODB DEFAULT CHARSET=utf8;
