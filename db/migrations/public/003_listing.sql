CREATE TABLE IF NOT EXISTS public.listing
(
    id serial NOT NULL,
    title text NOT NULL,
    description text,
    image_link text,
    price int NOT NULL,
    date_created timestamp without timezone NOT NULL,
    CONSTRAINT price CHECK (price >= 0), 
    CONSTRAINT listing_pkey PRIMARY KEY (id)
);

---- create above / drop below ----

DROP TABLE IF EXISTS public.listing;
