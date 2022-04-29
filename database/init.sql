-- create database gormyt_db
create table product
(
    id varchar(80)   not null
        primary key,
    name varchar(80)   not null,
    price numeric(8, 2) not null
);
/*
create table cliente
(
    id        serial
        primary key,
    nombre    varchar(80) not null,
    direccion varchar(80)
);
create table empleado
(
    id       serial
        primary key,
    nombre   varchar(80) not null,
    telefono varchar(15) not null
);

create table venta
(
    id         serial
        primary key,
    fecha      timestamp,
    cliente_id integer not null
        references cliente,
    empl_id    integer not null
        references empleado
);
create table detalle_venta
(
    venta_id        integer     not null
        references venta,
    producto_codigo varchar(80) not null
        references producto,
    primary key (venta_id, producto_codigo)
);
*/