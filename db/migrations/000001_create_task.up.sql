CREATE TABLE IF NOT EXISTS tasks(
    id serial PRIMARY KEY, 
    "name" VARCHAR(100) NOT NULL, 
    "desc" TEXT, 
    create_date TIMESTAMP NOT NULL
);


INSERT INTO tasks("name", "desc", create_date)
VALUES ('Создание отчёта', 'Подготовка и отправка отчёта о проделанной работе', '2025-02-24 09:00:00'),
('Исправление ошибки', 'Исправление критической ошибки в модуле авторизации', '2025-02-24 10:30:00'),
('Обновление документации', 'Обновление документации по проекту', '2025-02-24 11:15:00'),
('Планирование встречи', 'Организация планерки команды', '2025-02-24 14:00:00');
