
create table deviceMaster
(
	id int not null AUTO_INCREMENT PRIMARY key,
	device varchar(100),
	value decimal(8,4),
	updated_at datetime
)