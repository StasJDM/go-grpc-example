CREATE TABLE IF NOT EXISTS public.posts
(
    id         bigserial    NOT NULL,
    created_at timestamptz  NULL,
    updated_at timestamptz  NULL,
    deleted_at timestamptz  NULL,
    title      varchar(255) NOT NULL,
    content    text         NOT NULL,
    CONSTRAINT posts_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_users_deleted_at ON public.posts USING btree (deleted_at);