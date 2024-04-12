-- Create "users" table
CREATE TABLE "public"."users" (
  "user_id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "username" text NULL,
  "is_anonymous" boolean NULL DEFAULT false,
  "is_active" boolean NULL DEFAULT true,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("user_id")
);
-- Create "program_languages" table
CREATE TABLE "public"."program_languages" (
  "program_language_id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "name" text NULL,
  PRIMARY KEY ("program_language_id")
);
-- Create "code_snippets" table
CREATE TABLE "public"."code_snippets" (
  "code_snippet_id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "user_id" uuid NULL,
  "program_language_id" uuid NULL,
  "text" text NULL,
  "is_private" boolean NULL DEFAULT false,
  "is_archived" boolean NULL DEFAULT false,
  "is_draft" boolean NULL DEFAULT false,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("code_snippet_id"),
  CONSTRAINT "fk_code_snippets_program_language" FOREIGN KEY ("program_language_id") REFERENCES "public"."program_languages" ("program_language_id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_code_snippets_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("user_id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "code_snippet_ratings" table
CREATE TABLE "public"."code_snippet_ratings" (
  "code_snippet_rating_id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "code_snippet_id" uuid NULL,
  "user_id" uuid NULL,
  "rating" smallint NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("code_snippet_rating_id"),
  CONSTRAINT "fk_code_snippet_ratings_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("user_id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_code_snippets_code_snippet_ratings" FOREIGN KEY ("code_snippet_id") REFERENCES "public"."code_snippets" ("code_snippet_id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "review_comments" table
CREATE TABLE "public"."review_comments" (
  "comment_id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "user_id" uuid NULL,
  "code_snippet_id" uuid NULL,
  "reply_comment_id" text NULL,
  "text" text NULL,
  "line" bigint NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("comment_id"),
  CONSTRAINT "fk_code_snippets_review_comments" FOREIGN KEY ("code_snippet_id") REFERENCES "public"."code_snippets" ("code_snippet_id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_review_comments_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("user_id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
