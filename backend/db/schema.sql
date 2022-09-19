--
-- Name: test_table; Type: TYPE; Schema: testschema; Owner: tolga
--

DROP TABLE IF EXISTS hello_world;
CREATE TABLE hello_world (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL
);

