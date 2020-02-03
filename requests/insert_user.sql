INSERT INTO path (id, parent_id, created_at, accessed_at, modified_at) VALUES (?, NULL, ?, ?, ?);
INSERT INTO "group" (id, name) VALUES (?, ?);
INSERT INTO user (id, user_name, hash, storage_key, root_path_id, main_group_id)
 VALUES (?, ?, ?, ?, ?, ?);