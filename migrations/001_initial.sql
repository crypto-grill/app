-- migrate:up
SET
  TIME ZONE 'UTC';

CREATE TABLE IF NOT EXISTS user_ (
  id BIGINT PRIMARY KEY,
  pub_key VARCHAR(1024) NOT NULL
);

CREATE TABLE IF NOT EXISTS channel (
  id BIGINT PRIMARY KEY,
  sender_id BIGINT REFERENCES user_ (id) ON DELETE CASCADE,
  name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS subscriber (
  id SERIAL,
  user_id BIGINT REFERENCES user_ (id) ON DELETE CASCADE,
  channel_id BIGINT REFERENCES channel (id) ON DELETE CASCADE,
  PRIMARY KEY (user_id, channel_id)
);

CREATE TABLE IF NOT EXISTS subscribed_channel (
  id BIGINT PRIMARY KEY,
  channel_id BIGINT REFERENCES channel (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS message (
  id BIGINT NOT NULL,
  channel_id BIGINT REFERENCES channel (id) ON DELETE CASCADE,
  message VARCHAR(1024) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  PRIMARY KEY (id, channel_id)
);

CREATE TABLE IF NOT EXISTS subscription_proof (
  id BIGINT PRIMARY KEY,
  channel_id BIGINT REFERENCES channel (id) ON DELETE CASCADE,
  signature VARCHAR(1024) NOT NULL,
  message VARCHAR(1024) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- migrate:down
DROP TABLE IF EXISTS subscription_proof;

DROP TABLE IF EXISTS message;

DROP TABLE IF EXISTS subscribed_channel;

DROP TABLE IF EXISTS subscriber;

DROP TABLE IF EXISTS channel;

DROP TABLE IF EXISTS user_;
