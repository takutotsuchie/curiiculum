SET check_function_bodies = false;
CREATE SCHEMA "default";
CREATE TABLE "default".labels (
    id uuid NOT NULL,
    name text NOT NULL
);
CREATE TABLE "default".task_label_relations (
    id uuid NOT NULL,
    task_id uuid NOT NULL,
    label_id uuid NOT NULL
);
CREATE TABLE "default".tasks (
    id uuid NOT NULL,
    title text NOT NULL,
    explanation text,
    "limit" timestamp with time zone NOT NULL,
    priority integer NOT NULL,
    status text NOT NULL,
    user_id uuid NOT NULL
);
CREATE TABLE "default".test_migration (
    id integer NOT NULL
);
CREATE SEQUENCE "default".test_migration_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER SEQUENCE "default".test_migration_id_seq OWNED BY "default".test_migration.id;
CREATE TABLE "default".users (
    id uuid NOT NULL,
    name text NOT NULL
);
ALTER TABLE ONLY "default".test_migration ALTER COLUMN id SET DEFAULT nextval('"default".test_migration_id_seq'::regclass);
ALTER TABLE ONLY "default".labels
    ADD CONSTRAINT labels_pkey PRIMARY KEY (id);
ALTER TABLE ONLY "default".task_label_relations
    ADD CONSTRAINT task_label_relations_pkey PRIMARY KEY (id);
ALTER TABLE ONLY "default".tasks
    ADD CONSTRAINT tasks_pkey PRIMARY KEY (id);
ALTER TABLE ONLY "default".test_migration
    ADD CONSTRAINT test_migration_pkey PRIMARY KEY (id);
ALTER TABLE ONLY "default".users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);
