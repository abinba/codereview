-- Create "users" table
CREATE TABLE "public"."users" (
  "user_id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "username" text NOT NULL UNIQUE,
  "password" text NOT NULL,
  "is_active" boolean NULL DEFAULT true,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("user_id")
);
-- Create "program_languages" table
CREATE TABLE "public"."program_languages" (
  "program_language_id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "name" text NOT NULL,
  PRIMARY KEY ("program_language_id")
);
-- Create "code_snippets" table
CREATE TABLE "public"."code_snippets" (
  "code_snippet_id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "user_id" uuid NULL,
  "title" text NOT NULL,
  "is_private" boolean NULL DEFAULT false,
  "is_archived" boolean NULL DEFAULT false,
  "is_draft" boolean NULL DEFAULT false,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("code_snippet_id"),
  CONSTRAINT "fk_code_snippets_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("user_id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "code_snippet_versions" table
CREATE TABLE "public"."code_snippet_versions" (
  "code_snippet_version_id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "code_snippet_id" uuid NOT NULL,
  "program_language_id" uuid NOT NULL,
  "text" text NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("code_snippet_version_id"),
  CONSTRAINT "fk_code_snippet_versions_program_language" FOREIGN KEY ("program_language_id") REFERENCES "public"."program_languages" ("program_language_id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_code_snippets_code_snippet_versions" FOREIGN KEY ("code_snippet_id") REFERENCES "public"."code_snippets" ("code_snippet_id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create "code_snippet_ratings" table
CREATE TABLE "public"."code_snippet_ratings" (
  "code_snippet_rating_id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "code_snippet_version_id" uuid NOT NULL,
  "user_id" uuid NOT NULL,
  "rating" smallint NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("code_snippet_rating_id"),
  CONSTRAINT "fk_code_snippet_ratings_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("user_id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_code_snippet_versions_code_snippet_ratings" FOREIGN KEY ("code_snippet_version_id") REFERENCES "public"."code_snippet_versions" ("code_snippet_version_id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "notifications" table
CREATE TABLE "public"."notifications" (
  "notification_id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "notification_type" text NOT NULL,
  "user_id" uuid NOT NULL,
  "text" text NOT NULL,
  "created_at" timestamptz NULL,
  PRIMARY KEY ("notification_id"),
  CONSTRAINT "fk_notifications_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("user_id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "review_comments" table
CREATE TABLE "public"."review_comments" (
  "comment_id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "user_id" uuid NOT NULL,
  "code_snippet_version_id" uuid NOT NULL,
  "reply_comment_id" uuid NULL,
  "text" text NOT NULL,
  "line" bigint NULL,
  "is_generated" boolean NULL DEFAULT false,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("comment_id"),
  CONSTRAINT "fk_code_snippet_versions_review_comments" FOREIGN KEY ("code_snippet_version_id") REFERENCES "public"."code_snippet_versions" ("code_snippet_version_id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_review_comments_reply_comment" FOREIGN KEY ("reply_comment_id") REFERENCES "public"."review_comments" ("comment_id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_review_comments_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("user_id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Insert most popular programming languages
INSERT INTO public.program_languages (name) VALUES ('JavaScript');
INSERT INTO public.program_languages (name) VALUES ('Python');
INSERT INTO public.program_languages (name) VALUES ('Java');
INSERT INTO public.program_languages (name) VALUES ('C#');
INSERT INTO public.program_languages (name) VALUES ('C++');
INSERT INTO public.program_languages (name) VALUES ('TypeScript');
INSERT INTO public.program_languages (name) VALUES ('PHP');
INSERT INTO public.program_languages (name) VALUES ('Ruby');
INSERT INTO public.program_languages (name) VALUES ('Swift');
INSERT INTO public.program_languages (name) VALUES ('Go');
INSERT INTO public.program_languages (name) VALUES ('Kotlin');
INSERT INTO public.program_languages (name) VALUES ('Rust');
INSERT INTO public.program_languages (name) VALUES ('R');
INSERT INTO public.program_languages (name) VALUES ('Scala');
