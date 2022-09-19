
CREATE SCHEMA IF NOT EXISTS test_service;

--
-- Name: hello_world; Type: TYPE; Schema: test_service; Owner: tolga
--

DROP TABLE IF EXISTS test_service.hello_world;
CREATE TABLE test_service.hello_world (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    last_name VARCHAR NOT NULL
);
