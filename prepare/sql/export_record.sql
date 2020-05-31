CREATE TABLE `export_record` (
	id INT PRIMARY KEY,
	`staff_name` VARCHAR(32) NOT NULL,
	`time` DATETIME NOT NULL,
	`c_id` VARCHAR(32) NOT NULL,
	`batch_flag` VARCHAR(32) NOT NULL
)