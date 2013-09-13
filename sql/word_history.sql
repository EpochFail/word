create table word_history
(
  word_history_id serial NOT NULL,
  word_id integer,
  created_at timestamp without time zone default current_timestamp,
  constraint word_history_pkey primary key (word_history_id)
);
