CREATE TYPE role AS ENUM (
  'admin',
  'customer'
);

CREATE TYPE oderenum AS ENUM (
  'pending',
  'process',
  'done'
);

CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "username" varchar NOT NULL UNIQUE,
  "password" varchar NOT NULL,
  "email" varchar NOT NULL,
  "role" role NOT NULL,
  "full_name" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "order_items" (
  "order_id" int NOT NULL,
  "book_id" int NOT NULL,
  "quantity" int NOT NULL DEFAULT 1,
  "price_invoice" int NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "orders" (
  "id" SERIAL PRIMARY KEY,
  "user_id" int NOT NULL,
  "status" oderenum NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "book_categories" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "books" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar NOT NULL,
  "price" int NOT NULL,
  "quantity" int NOT NULL,
  "book_category_id" int NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

ALTER TABLE "books" ADD FOREIGN KEY ("book_category_id") REFERENCES "book_categories"("id") ON DELETE CASCADE;

ALTER TABLE "order_items" ADD PRIMARY KEY ("order_id", "book_id");

ALTER TABLE "order_items" ADD FOREIGN KEY ("order_id") REFERENCES "orders"("id") ON DELETE CASCADE;

ALTER TABLE "order_items" ADD FOREIGN KEY ("book_id") REFERENCES "books"("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("user_id") REFERENCES "users"("id") ON DELETE CASCADE;