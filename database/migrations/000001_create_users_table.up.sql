CREATE TABLE IF NOT EXISTS public.users(
   id SERIAL PRIMARY KEY,
   name text,
   login text UNIQUE,
   password text,
   registered_at timestamp with time zone DEFAULT now(),
   expired_at timestamp with time zone DEFAULT now() + interval '2min'
);