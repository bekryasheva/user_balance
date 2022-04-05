CREATE TABLE IF NOT EXISTS users (
    id bigint PRIMARY KEY,
    balance numeric
);

CREATE TABlE IF NOT EXISTS transactions (
  id bigint PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  user_id bigint REFERENCES users(id),
  type varchar(255),
  date timestamp,
  amount numeric,
  from_id bigint,
  to_id bigint
);

