--
-- Name: test_table; Type: TYPE; Schema: testschema; Owner: tolga
--

DROP TABLE IF EXISTS test_table;
CREATE TABLE test_table (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL
);

