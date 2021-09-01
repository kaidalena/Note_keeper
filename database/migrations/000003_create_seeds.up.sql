insert into public.users (id, name, login, password) values
(1, 'Tester', 'cat', '9003d1df22eb4d3820015070385194c8'),                   -- 'pwd' = '9003d1df22eb4d3820015070385194c8'
(2, 'Irina', 'dog', '9003d1df22eb4d3820015070385194c8'),
(3, 'Kaida Lena', 'little_coon', '9003d1df22eb4d3820015070385194c8'),
(4, 'Eva', 'fox', '9003d1df22eb4d3820015070385194c8');

insert into public.notes (user_id, text_note) values
(1, 'first Tester note'),
(1, 'second Tester note'),
(2, 'first Irina note'),
(2, 'second Irina note'),
(3, 'first Lena note'),
(3, 'second little_coon note'),
(4, 'first Eva note'),
(4, 'second Eva note');