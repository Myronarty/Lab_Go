CREATE TABLE IF NOT EXISTS kogut (
    id integer NOT NULL,
    name character varying(100) NOT NULL,
    age integer,
    sex boolean NOT NULL,
    CONSTRAINT kogut_age_check CHECK (((age >= 0) AND (age <= 30)))
);