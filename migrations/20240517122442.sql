-- Modify "users" table
ALTER TABLE "public"."users" DROP CONSTRAINT "users_username_key", ADD COLUMN "first_name" text NULL, ADD COLUMN "last_name" text NULL, ADD COLUMN "phone_number" text NULL;
