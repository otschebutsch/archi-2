DROP TABLE IF EXISTS "computers";
DROP TABLE IF EXISTS "disks";
CREATE TABLE "computers"
(
    "id_c" SERIAL PRIMARY KEY,
    "name" VARCHAR NOT NULL UNIQUE,
    "cpucount" INT,
    "totaldiskspace" BIGINT
);

CREATE TABLE "disks"
(
    "id_d" SERIAL PRIMARY KEY,
    "diskspace" BIGINT NOT NULL UNIQUE,
    "id_c" INTEGER REFERENCES computers (id_c)
);

-- Insert demo data.
INSERT INTO "computers" (id_c, name, cpucount, totaldiskspace) VALUES ('1', 'server-1', 4, '31457280');
INSERT INTO "computers" (id_c, name, cpucount, totaldiskspace) VALUES ('2', 'server-2', 6, '62914560');
INSERT INTO "disks" (id_d, diskspace) VALUES ('1', '31457280');
INSERT INTO "disks" (id_d, diskspace) VALUES ('2', '62914560');
