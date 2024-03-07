CREATE DATABASE user;

\c user

-- Execute SQL dump file
\i dockerfiles/user_dump.sql;