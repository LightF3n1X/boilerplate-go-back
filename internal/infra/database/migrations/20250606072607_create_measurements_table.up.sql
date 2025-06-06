CREATE TABLE IF NOT EXISTS public.measurements
(
    id SERIAL PRIMARY KEY,
    device_id INTEGER REFERENCES public.devices(id),
    room_id INTEGER REFERENCES public.rooms(id),
    Value TEXT NOT NULL,
    created_date    timestamptz NOT NULL,
    updated_date    timestamptz NOT NULL,
    deleted_date    timestamptz
);