create table cars
(
    id          varchar(255) not null primary key,
    make        varchar(255) not null,
    model       varchar(255) not null,
    description text null
);

create table pricing
(
    car_id  varchar(255) not null,
    mileage int          not null,
    term    int          not null,
    price   int          not null,

    constraint pricing_cars_id_fk foreign key (car_id) references cars (id) on update cascade on delete cascade
);

