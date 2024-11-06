-- name: GetMovies :many
select * from movie;

-- name: SaveMovie :exec
insert into movie (id, title, release_date)
values ($1, $2, $3)
on conflict(id)
do update set
  title = excluded.title,
  release_date = excluded.release_date;

-- name: DeleteMovie :exec
delete from movie
where id = $1;
