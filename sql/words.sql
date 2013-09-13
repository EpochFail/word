CREATE TABLE words
(
  word_id serial NOT NULL,
  word character varying,
  rating integer,
  CONSTRAINT words_pkey PRIMARY KEY (word_id )
);

CREATE INDEX word_id_index
  ON words (word_id );
