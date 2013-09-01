CREATE TABLE votes
(
  vote_id serial NOT NULL,
  word character varying,
  vote integer,
  created_at timestamp without time zone default current_timestamp, 
  CONSTRAINT votes_pkey PRIMARY KEY (vote_id)
);

CREATE INDEX vote_id_index
  on votes (vote_id);
