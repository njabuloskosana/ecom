CREATE OR REPLACE FUNCTION notify_listener(topic TEXT, payload JSONB)
RETURNS VOID AS $$
BEGIN
  PERFORM pg_notify(topic, payload::TEXT);
END;
$$ LANGUAGE plpgsql;


