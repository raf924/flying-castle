-- 1: base64 encoded key
-- 2: temp path to chunk 1
-- 3: temp path to chunk 2
-- 4: base64 encoded hash
-- 5: base64 encoded salt
INSERT INTO storage_key
VALUES (1, ?, current_timestamp);
INSERT INTO chunk
VALUES (1, ?, NULL);
INSERT INTO chunk
VALUES (2, ?, NULL);
UPDATE chunk
set next_chunk = 2
where id = 1;
INSERT INTO path(id, name, parent_id, created_at, accessed_at, modified_at)
VALUES (1, 'rafael', NULL, current_timestamp, current_timestamp, current_timestamp);
INSERT INTO path(id, name, parent_id, created_at, accessed_at, modified_at)
VALUES (2, 'file1', 1, current_timestamp, current_timestamp, current_timestamp);
INSERT INTO "group"(id, name) VALUES (1, 'rafael');
INSERT INTO folder(id, path_id) VALUES (1, 1);
INSERT INTO user(id, user_name, hash, salt, root_folder_id, main_group_id)
VALUES (1, 'rafael', ?, ?, 1, 1);
INSERT INTO file(id, first_chunk_id, path_id, size, modified_at)
VALUES (1, 1, 2, 10, current_timestamp);