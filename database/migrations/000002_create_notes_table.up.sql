CREATE TABLE IF NOT EXISTS public.notes(
   id SERIAL PRIMARY KEY,
   user_id integer REFERENCES public.users,
   text_note text
);