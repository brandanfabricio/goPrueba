create database inventario;

use inventario;

create Table users (
  id not null auto_increment,
  email varchar(255) not null,
  name varchar(255) not null,
  password varchar(255) not null
  primary key(id)
);

create table product (
  id int not null auto_increment,
  name varchar(255) not null,
  description varchar(255) not null,
  price float not null,
  user_id int not null,
  primary key (id),
  foreign key (user_id) references users(id) 
);
create table roles(
   id int not null,
  name varchar(255) not null,
  primary key (id),
);

create table users_role(
  id int not null auto_increment,
  user_id int not null,
  role_id int not null,
  primary key (id),
  foreign key (user_id) references users(id),
  foreign key (role_id) references users(roles) 

);