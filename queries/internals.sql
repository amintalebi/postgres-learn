CREATE function table_name_to_create_table(table_name text)
returns text as $$
select 'CREATE TABEL ' || relname || '(' || string_agg(attname || ' ' || typname, ', ') || ')'
from pg_class, pg_attribute, pg_type
where relname=table_name and pg_class.oid= pg_attribute.attrelid and attnum > 0 and not attisdropped and pg_type.oid=atttypid
group by relname;
$$ language sql;