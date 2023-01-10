#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
create table tasks (
  id serial not null unique,
  text varchar(100),
  day varchar(64),
  reminder boolean,
  primary key(id)
);


insert into tasks(text, day,reminder)
values
    ('Doctors Appointment', 'March 5th at 2:30pm',TRUE),
    ('Meeting with boss', 'March 6th at 1:30pm',TRUE),
    ('Food shopping', 'March 7th at 2:00pm',FALSE);
EOSQL
