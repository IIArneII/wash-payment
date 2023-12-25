-- +goose Up
-- +goose StatementBegin

ALTER TABLE public.organizations DROP constraint organizations_version_check;
ALTER TABLE public.organizations ADD constraint new_organizations_version_check CHECK((version>=0));

ALTER TABLE public.groups DROP constraint groups_version_check;
ALTER TABLE public.groups ADD constraint new_groups_version_check CHECK((version>=0));

ALTER TABLE public.users DROP constraint users_version_check;
ALTER TABLE public.users ADD constraint new_users_version_check CHECK((version>=0));

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin



-- +goose StatementEnd
