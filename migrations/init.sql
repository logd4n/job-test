-- Проверяем существование базы данных
DO $$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_database WHERE datname = 'test') THEN
        CREATE DATABASE test;
    END IF;
END
$$;

-- Проверяем существование пользователя
DO $$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_user WHERE usename = 'postgres') THEN
        CREATE USER myuser WITH ENCRYPTED PASSWORD '12345';
    END IF;
END
$$;

-- Предоставляем права
GRANT ALL PRIVILEGES ON DATABASE test TO postgres;