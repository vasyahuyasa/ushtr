CREATE SCHEMA ushtr_1;
CREATE SEQUENCE ushtr_1.urls_id_seq;
CREATE OR REPLACE FUNCTION ushtr_1.next_id(OUT result bigint) AS $$
DECLARE
    shard_id int := 1;
BEGIN
    result := shard_id;
    result := result | (nextval('ushtr_1.urls_id_seq') << 10);
END;
    $$ LANGUAGE PLPGSQL;

CREATE TABLE ushtr_1.urls (
    "id" bigint NOT NULL DEFAULT ushtr_1.next_id(),
	"url" text NOT NULL,
	"created_at" timestamp without time zone default (now() at time zone 'utc')
  )
