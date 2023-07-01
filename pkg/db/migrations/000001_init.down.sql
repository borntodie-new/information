-- 删除 user_fans 表的外键
ALTER TABLE user_fans DROP CONSTRAINT user_fans_followed_id_fkey;
ALTER TABLE user_fans DROP CONSTRAINT user_fans_follower_id_fkey;

-- 删除 news 表的外键
ALTER TABLE news DROP CONSTRAINT news_user_id_fkey;
ALTER TABLE news DROP CONSTRAINT news_category_id_fkey;

-- 删除 collections 表的外键
ALTER TABLE collections DROP CONSTRAINT collections_user_id_fkey;
ALTER TABLE collections DROP CONSTRAINT collections_news_id_fkey;

-- 删除 comments 表的外键
ALTER TABLE comments DROP CONSTRAINT comments_user_id_fkey;
ALTER TABLE comments DROP CONSTRAINT comments_parent_id_fkey;
ALTER TABLE comments DROP CONSTRAINT comments_news_id_fkey;

-- 删除 comment_likes 表的外键
ALTER TABLE comment_likes DROP CONSTRAINT comment_likes_user_id_fkey;
ALTER TABLE comment_likes DROP CONSTRAINT comment_likes_comment_id_fkey;


DROP TABLE IF EXISTS "users";
DROP TABLE IF EXISTS "news";
DROP TABLE IF EXISTS "comments";
DROP TABLE IF EXISTS "comment_likes";
DROP TABLE IF EXISTS "collections";
DROP TABLE IF EXISTS "user_fans";
DROP TABLE IF EXISTS "categories";
