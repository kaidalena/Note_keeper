CREATE TABLE IF NOT EXISTS public.users(
   id SERIAL PRIMARY KEY,
   name text,
   login text UNIQUE,
   password text,
   expired_at timestamp with time zone
);