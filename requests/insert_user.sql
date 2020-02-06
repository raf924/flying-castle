INSERT INTO "group" (id, name) VALUES (?, ?);
INSERT INTO user (id, user_name, hash, salt, root_folder_id, main_group_id)
 VALUES (?, ?, ?, ?, ?, ?);