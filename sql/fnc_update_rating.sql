create or replace function fnc_update_rating() returns trigger as $$
begin
  update words set rating = rating + NEW.vote where word = NEW.word;
  return new;
end;
$$ language plpgsql;

create trigger trg_update_rating after insert on votes
for each row execute procedure fnc_update_rating();
