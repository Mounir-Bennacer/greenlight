#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
	      CREATE EXTENSION IF NOT EXISTS citext;
		    -- # CREATE ROLE green WITH PASSWORD 'light' LOGIN;

		    -- # GRANT ALL PRIVILEGES ON SCHEMA public TO green;
		    -- # GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO green;
		    -- # GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO green;

		    -- # ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL ON TABLES TO green;
		    -- # ALTER DATABASE greenlight OWNER TO green;
EOSQL
