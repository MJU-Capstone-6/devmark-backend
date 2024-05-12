-- Create "workspace_user_category" view
CREATE VIEW "public"."workspace_user_category" ("id", "name", "created_at", "updated_at", "categories", "users") AS SELECT w.id,
    w.name,
    w.created_at,
    w.updated_at,
    json_agg(c.*) AS categories,
    json_agg(u.*) AS users
   FROM ((((workspace w
     JOIN workspace_category wc ON ((w.id = wc.workspace_id)))
     JOIN category c ON ((wc.category_id = c.id)))
     JOIN workspace_user wu ON ((w.id = wu.workspace_id)))
     JOIN "user" u ON ((wu.user_id = u.id)))
  GROUP BY w.id;
