-- Modify "workspace_user_category" view
CREATE OR REPLACE VIEW "public"."workspace_user_category" ("id", "name", "created_at", "updated_at", "categories", "users") AS SELECT w.id,
    w.name,
    w.created_at,
    w.updated_at,
    json_agg(DISTINCT c.*) AS categories,
    json_agg(DISTINCT u.*) AS users
   FROM ((((workspace w
     LEFT JOIN workspace_category wc ON ((w.id = wc.workspace_id)))
     LEFT JOIN category c ON ((wc.category_id = c.id)))
     LEFT JOIN workspace_user wu ON ((w.id = wu.workspace_id)))
     LEFT JOIN "user" u ON ((wu.user_id = u.id)))
  GROUP BY w.id;
