CREATE TABLE IF NOT EXISTS public.devices
(
    id SERIAL PRIMARY KEY,
    house_id INTEGER REFERENCES public.houses(id),
    room_id INTEGER REFERENCES public.rooms(id),
    uuid TEXT NOT NULL,
    serial_numbers TEXT NOT NULL,
    characteristics TEXT NOT NULL,
    category TEXT NOT NULL,
    Units TEXT NOT NULL,
    power_consumption INT NOT NULL,
    created_date    timestamptz NOT NULL,
    updated_date    timestamptz NOT NULL,
    deleted_date    timestamptz
);