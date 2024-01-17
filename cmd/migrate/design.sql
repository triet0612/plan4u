create table kanban_column (
    colname text,
    color text,
    note text,
    col_limit integer
);
-- status == colname
create table kanban_item (
    title text,
    detail text,
    item_status text,
    deadline date
);
create table schedule_item (
    item_id integer,
    note text,
    starttime text,
    endtime text
);
create table calender_item (
    title text,
    detail text,
    item_date date,
    color text
);
