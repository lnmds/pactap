-- Which version is this repository on?
create table if not exists build (
    build int
);

-- Current state of packages
create table if not exists packages (
    name string,
    version string,
    build int
);

create table if not exists depedencies (
    package string references packages (name),
    depedency string references packages (name)
);

-- From this you generate difference files to send
-- to users as a repo patch file
create table if not exists logs (
    version int references versioning (version)
    -- and then what ?? w ?
    -- maybe something serializable like json??? idk
);

