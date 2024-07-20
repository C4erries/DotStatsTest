CREATE TABLE users(
    id bigserial not null primary key,
    nickname varchar not null,
    email varchar not null unique,
    encrypted_password varchar not null,
    player_id bigint not null unique
);

CREATE TABLE matches(
    match_id bigint not null unique primary key,
    team_R bigint[5],
    team_D bigint[5],
    time_length time not null,
    result varchar
);

CREATE TABLE stats(
    player_id bigint unique not null primary key ,
    player_stats jsonb,
    heroes_stats jsonb,
    matches_stats jsonb,
    wordcloud_stats jsonb,
    FOREIGN KEY (player_id) REFERENCES users (player_id) ON DELETE CASCADE
);