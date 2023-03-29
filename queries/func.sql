create or replace function update_views_2()
  returns trigger
  language plpgsql
  as
  $$
  begin
   refresh materialized view top_authors;
   return NEW;
  end;
  $$;

create trigger top_author_update after update on book_reviews execute function update_views_2();