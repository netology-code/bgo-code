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
