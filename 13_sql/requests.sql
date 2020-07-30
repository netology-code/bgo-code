-- если указываем только имя таблицы, то обязаны заполнять все поля
INSERT INTO clients
VALUES (1,
        'petr',
        'password-hash',
        'Пётр Николаевич Иванов',
        '8204 95523',
        '1970.01.30',
        'ACTIVE',
        NOW());

SELECT *
FROM clients;

DELETE
FROM clients
WHERE id = 1;

-- если указываем имя таблицы и столбцы,
-- то обязаны заполнять только то, что указали
-- остальное примет значение по умолчанию либо NULL
-- (если не установлено NOT NULL)
INSERT INTO clients(login, password, full_name, passport, birthday, status)
VALUES ('petr',
        'password-hash',
        'Пётр Николаевич Иванов',
        '8204 95523',
        '1970.01.30',
        'ACTIVE');


INSERT INTO clients(login, password, full_name, passport, birthday, status)
VALUES ('vasya', 'password-hash', 'Василий Николаевич Иванов',
        '8205 96563', '1970.01.30', 'ACTIVE'),
       ('masha', 'password-hash', 'Мария Ивановна Петрова',
        '8205 48839', '1990.11.21', 'ACTIVE'),
       ('dasha', 'password-hash', 'Дарья Ивановна Крылова',
        '8205 94483', '1995.04.27', 'ACTIVE')
;


