CREATE TABLE IF NOT EXISTS public.user_listing
(
    id_user serial NOT NULL,
    id_listing serial NOT NULL,
    CONSTRAINT user_listing_pkey PRIMARY KEY (id_user, id_listing),
    CONSTRAINT user_listing_id_user_fkey FOREIGN KEY (id_user)
        REFERENCES public.user (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
        NOT VALID,
    CONSTRAINT user_listing_id_listing_fkey FOREIGN KEY (id_listing)
        REFERENCES public.listing (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
        NOT VALID
)

---- create above / drop below ----

DROP TABLE IF EXISTS public.user_listing;
