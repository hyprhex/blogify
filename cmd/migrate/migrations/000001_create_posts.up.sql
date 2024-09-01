CREATE TABLE IF NOT EXISTS posts (
  id bigserial primary key,
  title text not null,
  content text not null,
  category text not null, 
  created_at timestamp(0) with time zone not null default now(),
  updated_at timestamp(0) with time zone not null default now()
);
