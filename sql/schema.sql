CREATE TABLE `users`
(
    id   bigint auto_increment,
    name varchar(255) NOT NULL,
    PRIMARY KEY (`id`)
);

INSERT INTO `users` (`name`)
VALUES ('Solomon'),
       ('Menelik');

CREATE TABLE `climbs`
(
    id varchar(255) NOT NULL,
    date varchar(255) NOT NULL,
    grade varchar(255) NOT NULL,
    description TEXT,
    area varchar(255) NOT NULL,
    PRIMARY KEY (`id`)
);

INSERT INTO `users` (`name`)
VALUES ('Solomon'),
       ('Menelik');