Create a mysql database and a table for this project.
create database inventory;
Query OK, 1 row affected (0.02 sec)

mysql> use inventory;
Database changed
mysql> create table products(
    -> id int NOT NULL AUTO_INCREMENT,
    -> name varchar(255) NOT NULL,
    -> quantity int,
    -> price float(10,7),
    -> PRIMARY KEY(id)
    -