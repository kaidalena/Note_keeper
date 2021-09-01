CREATE TABLE IF NOT EXISTS public.notes(
   id SERIAL PRIMARY KEY,
   user_id integer REFERENCES public.users,
   created_at timestamp with time zone DEFAULT now(),
   text_note text
);