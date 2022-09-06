
SET default_tablespace = '';
SET default_table_access_method = heap;

--
-- Name: go_template_service; Type: SCHEMA; Schema: -; Owner: tolga
--

CREATE SCHEMA testschema;
ALTER SCHEMA testschema OWNER TO tolga;

--
-- Name: test_table; Type: TYPE; Schema: testschema; Owner: tolga
--

CREATE TABLE testschema.test_table (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL
);

