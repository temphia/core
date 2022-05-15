create table example(
    __id serial primary key,
    __version integer not null default 0,
    __mod_sig text,
    w1 text not null,
    sh1 text,
    pn1 text,
    ss1 text,
    rf1 text,
    f1 text,
    mf1 text [],
    cb1 boolean,
    cy1 decimal,
    n1 integer,
    l1 geography(point, 4326),
    d1 timestamp
);