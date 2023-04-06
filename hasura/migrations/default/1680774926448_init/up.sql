SET check_function_bodies = false;
CREATE TABLE public.task_label_relations (
    id integer NOT NULL,
    task_id uuid NOT NULL,
    label_id uuid NOT NULL
);
CREATE SEQUENCE public.task_label_relations_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER SEQUENCE public.task_label_relations_id_seq OWNED BY public.task_label_relations.id;
CREATE TABLE public.tasks (
    id uuid NOT NULL,
    title text NOT NULL,
    explanation text,
    "limit" date NOT NULL,
    priority integer NOT NULL,
    status integer NOT NULL,
    user_id uuid NOT NULL
);
CREATE TABLE public.users (
    id uuid NOT NULL,
    name text NOT NULL
);
ALTER TABLE ONLY public.task_label_relations ALTER COLUMN id SET DEFAULT nextval('public.task_label_relations_id_seq'::regclass);
ALTER TABLE ONLY public.task_label_relations
    ADD CONSTRAINT task_label_relations_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.tasks
    ADD CONSTRAINT tasks_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.tasks
    ADD CONSTRAINT tasks_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
