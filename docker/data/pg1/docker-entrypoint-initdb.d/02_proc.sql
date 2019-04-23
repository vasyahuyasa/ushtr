CREATE OR REPLACE FUNCTION ushtr1.next_id(OUT result bigint) AS $$
DECLARE
    start_epoch bigint := 1555977600000;
    seq_id bigint;
    now_millis bigint;
    shard_id int := 5;
BEGIN
    SELECT nextval('ushtr1.table_id_seq') %% 1024 INTO seq_id;
    SELECT FLOOR(EXTRACT(EPOCH FROM clock_timestamp()) * 1000) INTO now_millis;
    result := (now_millis - start_epoch) << 23;
    result := result | (shard_id <<10);
    result := result | (seq_id);
END;
    $$ LANGUAGE PLPGSQL;
