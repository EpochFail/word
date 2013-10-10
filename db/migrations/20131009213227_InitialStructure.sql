
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE words
(
  word_id serial NOT NULL,
  word character varying,
  rating integer,
  CONSTRAINT words_pkey PRIMARY KEY (word_id )
);

CREATE INDEX word_id_index
  ON words (word_id );

create table word_history
(
  word_history_id serial NOT NULL,
  word_id integer,
  created_at timestamp without time zone default current_timestamp,
  constraint word_history_pkey primary key (word_history_id)
);

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

-- +goose StatementBegin
create or replace function fnc_update_rating() returns trigger as $$
begin
  update words set rating = rating + NEW.vote where word = NEW.word;
  return new;
end;
$$ language plpgsql;
-- +goose StatementEnd

create trigger trg_update_rating after insert on votes
for each row execute procedure fnc_update_rating();

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE words;

DROP TABLE word_history;

DROP TABLE votes;

DROP FUNCTION fnc_update_rating();
