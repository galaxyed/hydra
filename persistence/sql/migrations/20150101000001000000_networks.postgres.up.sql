-- Migration generated by the command below; DO NOT EDIT.
-- hydra:generate hydra migrations gen
CREATE TABLE "networks" (
  "id" UUID NOT NULL,
  PRIMARY KEY("id"),
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL
);

INSERT INTO networks (id, created_at, updated_at) VALUES ('24704dcb-0ab9-4bfa-a84c-405932ae53fe', '2013-10-07 08:23:19', '2013-10-07 08:23:19');
