CREATE TABLE IF NOT EXISTS public.user
(
    id serial NOT NULL,
    login text NOT NULL,
    password_hash text NOT NULL,
    CONSTRAINT user_pkey PRIMARY KEY (id)
);

---- create above / drop below ----

DROP TABLE IF EXISTS public.user;
