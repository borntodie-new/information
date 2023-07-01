CREATE TABLE "users"
(
    "id"            bigserial PRIMARY KEY,
    "nickname"      varchar        NOT NULL,
    "username"      varchar UNIQUE NOT NULL,
    "phone"         varchar UNIQUE NOT NULL,
    "email"         varchar UNIQUE NOT NULL,
    "avatar"        varchar        NOT NULL DEFAULT '/static/images/backend/default-avatar.png',
    "last_login_at" timestamptz    NOT NULL DEFAULT (000 - 01 - 01),
    "created_at"    timestamptz    NOT NULL DEFAULT (now()),
    "updated_at"    timestamptz    NOT NULL DEFAULT (000 - 01 - 01)
);

CREATE TABLE "user_fans"
(
    "id"          bigserial PRIMARY KEY,
    "follower_id" int         NOT NULL,
    "followed_id" int         NOT NULL,
    "created_at"  timestamptz NOT NULL DEFAULT (now()),
    "updated_at"  timestamptz NOT NULL DEFAULT (000 - 01 - 01)
);

CREATE TABLE "news"
(
    "id"          bigserial PRIMARY KEY,
    "title"       varchar     NOT NULL,
    "source"      varchar     NOT NULL,
    "digital"     varchar     NOT NULL,
    "content"     text        NOT NULL,
    "clicks"      int                  DEFAULT 0,
    "status"      int         NOT NULL DEFAULT 1,
    "reason"      text,
    "images"      text,
    "updated_at"  timestamptz NOT NULL DEFAULT (000 - 01 - 01),
    "created_at"  timestamptz NOT NULL DEFAULT (now()),
    "category_id" int         NOT NULL,
    "user_id"     int         NOT NULL
);

CREATE TABLE "categories"
(
    "id"         bigserial PRIMARY KEY,
    "name"       varchar     NOT NULL,
    "updated_at" timestamptz NOT NULL DEFAULT (000 - 01 - 01),
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "collections"
(
    "id"         bigserial PRIMARY KEY,
    "user_id"    int         NOT NULL,
    "news_id"    int         NOT NULL,
    "updated_at" timestamptz NOT NULL DEFAULT (000 - 01 - 01),
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "comments"
(
    "id"         bigserial PRIMARY KEY,
    "content"    varchar     NOT NULL,
    "likes"      int                  DEFAULT 0,
    "parent_id"  int         NOT NULL,
    "user_id"    int         NOT NULL,
    "news_id"    int         NOT NULL,
    "updated_at" timestamptz NOT NULL DEFAULT (000 - 01 - 01),
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "comment_likes"
(
    "id"         bigserial PRIMARY KEY,
    "comment_id" int         NOT NULL,
    "user_id"    int         NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (000 - 01 - 01)
);

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "users" ("email");

CREATE UNIQUE INDEX ON "user_fans" ("follower_id", "followed_id");

CREATE UNIQUE INDEX ON "collections" ("user_id", "news_id");

COMMENT
ON COLUMN "users"."nickname" IS '昵称';

COMMENT
ON COLUMN "users"."username" IS '用户名';

COMMENT
ON COLUMN "users"."phone" IS '手机号';

COMMENT
ON COLUMN "users"."email" IS '邮箱';

COMMENT
ON COLUMN "users"."avatar" IS '头像地址';

COMMENT
ON COLUMN "users"."last_login_at" IS '最后一次登陆时间';

COMMENT
ON COLUMN "users"."created_at" IS '创建时间';

COMMENT
ON COLUMN "users"."updated_at" IS '修改时间';

COMMENT
ON COLUMN "user_fans"."follower_id" IS '关注者ID';

COMMENT
ON COLUMN "user_fans"."followed_id" IS '被关注者ID';

COMMENT
ON COLUMN "user_fans"."created_at" IS '创建时间';

COMMENT
ON COLUMN "user_fans"."updated_at" IS '修改时间';

COMMENT
ON COLUMN "news"."title" IS '新闻标题';

COMMENT
ON COLUMN "news"."source" IS '新闻来源';

COMMENT
ON COLUMN "news"."digital" IS '新闻摘要';

COMMENT
ON COLUMN "news"."content" IS '新闻内容';

COMMENT
ON COLUMN "news"."clicks" IS '新闻浏览量';

COMMENT
ON COLUMN "news"."status" IS '0-通过,1-审核中,2-不通过';

COMMENT
ON COLUMN "news"."reason" IS '不通过的理由';

COMMENT
ON COLUMN "news"."images" IS '新闻图片,序列化后的串';

COMMENT
ON COLUMN "news"."updated_at" IS '修改时间';

COMMENT
ON COLUMN "news"."created_at" IS '创建时间';

COMMENT
ON COLUMN "categories"."name" IS '分类名称';

COMMENT
ON COLUMN "categories"."updated_at" IS '修改时间';

COMMENT
ON COLUMN "categories"."created_at" IS '创建时间';

COMMENT
ON COLUMN "collections"."user_id" IS '收藏者的用户ID';

COMMENT
ON COLUMN "collections"."news_id" IS '收藏的新闻ID';

COMMENT
ON COLUMN "collections"."updated_at" IS '修改时间';

COMMENT
ON COLUMN "collections"."created_at" IS '创建时间';

COMMENT
ON COLUMN "comments"."content" IS '评论内容';

COMMENT
ON COLUMN "comments"."likes" IS '点赞人数';

COMMENT
ON COLUMN "comments"."parent_id" IS '父评论ID';

COMMENT
ON COLUMN "comments"."user_id" IS '评论人ID';

COMMENT
ON COLUMN "comments"."news_id" IS '评论新闻ID';

COMMENT
ON COLUMN "comments"."updated_at" IS '修改时间';

COMMENT
ON COLUMN "comments"."created_at" IS '创建时间';

COMMENT
ON COLUMN "comment_likes"."comment_id" IS '点赞评论ID';

COMMENT
ON COLUMN "comment_likes"."user_id" IS '点赞用户ID';

COMMENT
ON COLUMN "comment_likes"."created_at" IS '创建时间';

COMMENT
ON COLUMN "comment_likes"."updated_at" IS '修改时间';

ALTER TABLE "user_fans"
    ADD FOREIGN KEY ("follower_id") REFERENCES "users" ("id");

ALTER TABLE "user_fans"
    ADD FOREIGN KEY ("followed_id") REFERENCES "users" ("id");

ALTER TABLE "news"
    ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "news"
    ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "collections"
    ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "collections"
    ADD FOREIGN KEY ("news_id") REFERENCES "news" ("id");

ALTER TABLE "comments"
    ADD FOREIGN KEY ("parent_id") REFERENCES "comments" ("id");

ALTER TABLE "comments"
    ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "comments"
    ADD FOREIGN KEY ("news_id") REFERENCES "news" ("id");

ALTER TABLE "comment_likes"
    ADD FOREIGN KEY ("comment_id") REFERENCES "comments" ("id");

ALTER TABLE "comment_likes"
    ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
