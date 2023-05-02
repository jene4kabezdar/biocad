CREATE TABLE messages (
    n bigserial not null primary key,
    mqtt varchar,
    invid varchar not null default '',
    unit_guid varchar not null default '',
    msg_id varchar not null default '',
    text varchar not null default '',
    context varchar,
    class varchar not null default '',
    level int not null default 100,
    area varchar not null default 'LOCAL',
    addr varchar not null default '',
    block varchar not null default '',
    type varchar,
    bit varchar,
    invert_bit varchar 
);