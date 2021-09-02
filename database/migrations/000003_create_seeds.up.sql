DO $$
BEGIN
    IF (SELECT count(id) FROM public.users) > 0 THEN
        DELETE FROM public.users;
    END IF;

    IF (SELECT count(id) FROM public.notes) > 0 THEN
        DELETE FROM public.notes;
    END IF;
END $$;

INSERT INTO public.users (id, name, login, password) VALUES
    (1, 'Tester', 'cat', '9003d1df22eb4d3820015070385194c8'),                   -- 'pwd' = '9003d1df22eb4d3820015070385194c8'
    (2, 'Irina', 'dog', '9003d1df22eb4d3820015070385194c8'),
    (3, 'Kaida Lena', 'little_coon', '9003d1df22eb4d3820015070385194c8'),
    (4, 'Eva', 'fox', '9003d1df22eb4d3820015070385194c8');


INSERT INTO public.notes (user_id, text_note, created_at) VALUES
    (1, 'first Tester note', now() + interval '10min'),
    (1, 'second Tester note', now() + interval '5min'),
    (2, 'first Irina note', now() + interval '1min'),
    (2, 'second Irina note', now() + interval '8min'),
    (3, 'first Lena note', now() + interval '1min'),
    (3, 'second little_coon note', now() + interval '2min'),
    (4, 'first Eva note', now() + interval '7min'),
    (4, 'second Eva note', now() + interval '3min');