select * from employees;

explain select * from employees where last_name='wang' and first_name LIKE '%zi';

set optimizer_switch='index_condition_pushdown=off';



CREATE TABLE `employees`( 
    `emp_no`int(11)NOT NULL,
    `birth_date`date NULL,
    `first_name`varchar(14)NOT NULL,
    `last_name`varchar(16)NOT NULL, 
    `gender`enum('M','F')NOT NULL, 
    `hire_date`date NULL, 
    PRIMARY KEY (`emp_no`)
)ENGINE=InnoDB DEFAULT CHARSET=latin1;

alter table employees add index idx_lastname_firstname(last_name,first_name); -- 创建联合索引

INSERT INTO `employees`(`emp_no`,`birth_date`,`first_name`,`last_name`,`gender`,`hire_date`)VALUES(1, NULL,'698','liu','F',NULL); 
INSERT INTO `employees`(`emp_no`,`birth_date`,`first_name`,`last_name`,`gender`,`hire_date`)VALUES(2, NULL,'d99','zheng','F',NULL); 
INSERT INTO `employees`(`emp_no`,`birth_date`,`first_name`,`last_name`,`gender`,`hire_date`)VALUES(3, NULL,'e08','huang','F',NULL); 
INSERT INTO `employees`(`emp_no`,`birth_date`,`first_name`,`last_name`,`gender`,`hire_date`)VALUES(4, NULL,'59d','lu','F',NULL); 
INSERT INTO `employees`(`emp_no`,`birth_date`,`first_name`,`last_name`,`gender`,`hire_date`)VALUES(5, NULL,'0dc','yu','F',NULL); 
INSERT INTO`employees`(`emp_no`,`birth_date`,`first_name`,`last_name`,`gender`,`hire_date`)VALUES(6, NULL,'989','wang','F',NULL); 
INSERT INTO`employees`(`emp_no`,`birth_date`,`first_name`,`last_name`,`gender`,`hire_date`)VALUES(7, NULL,'e38','wang','F',NULL); 
INSERT INTO`employees`(`emp_no`,`birth_date`,`first_name`,`last_name`,`gender`,`hire_date`)VALUES(8, NULL,'0zi','wang','F',NULL); 
INSERT INTO`employees`(`emp_no`,`birth_date`,`first_name`,`last_name`,`gender`,`hire_date`)VALUES(9, NULL,'dc9','xie','F',NULL); 
INSERT INTO `employees`(`emp_no`,`birth_date`,`first_name`,`last_name`,`gender`,`hire_date`)VALUES(10, NULL,'5ba','zhou','F',NULL);
