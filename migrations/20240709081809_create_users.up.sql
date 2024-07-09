CREATE TABLE users(
    id bigserial not null primary key,
    nickname varchar not null,
    email varchar not null unique,
    encrypted_password varchar not null,
    player_id bigint not null unique
);

CREATE TABLE matches(
    match_id bigint not null unique,
    team_R bigint[5],
    team_D bigint[5],
    time_length time not null,
    result varchar
);