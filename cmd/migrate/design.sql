create table kanban_column (
    colname text,
    color text,
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
insert into kanban_column values
("Todo", "#6b1d32", 6);
insert into kanban_column values
("In progress", "#8c8120", 6);
insert into kanban_column values
("Done", "#1a6220", 6);
insert into kanban_column values
("Staging", "#7a69d1", 6);
insert into kanban_column values
("Deploy", "#fafcfe", 6);

insert into kanban_item values
("Task 1", "implement kanban board", "Done", "2024-20-01T10:00:00");
insert into kanban_item values
("Task 2", "implement calendar", "Todo", "2024-20-01T10:00:00");
insert into kanban_item values
("Task 4", "implement schedule", "In progress", "2024-20-01T10:00:00");
insert into kanban_item values
("Task 3", "implement controls", "In progress", "2024-20-01T10:00:00");
