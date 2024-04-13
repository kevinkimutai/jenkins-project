-- name: GetMovie :one
SELECT movies.*,COALESCE(AVG(ratings.rating), 0.0) AS average_rating
FROM movies
LEFT JOIN ratings 
ON movies.id = ratings.movie_id
WHERE movies.id = $1 
GROUP BY movies.id
LIMIT 1;

-- -- name: ListMovies :many
-- SELECT movies.*,AVG(ratings.rating) AS average_rating
-- FROM movies
-- LEFT JOIN ratings 
-- ON movies.id = ratings.movie_id;


-- name: ListMovies :many
SELECT movies.*,COALESCE(AVG(ratings.rating), 0.0) AS average_rating
FROM movies
LEFT JOIN ratings 
ON movies.id = ratings.movie_id
WHERE (movies.title ILIKE '%' || $1 || '%' OR $1 IS NULL)
  AND (movies.release_date >= $2 OR $2 IS NULL)
  AND (movies.release_date <= $3 OR $3 IS NULL)
GROUP BY movies.id
ORDER BY movies.id 
LIMIT $4 OFFSET $5;

-- name: CountMovies :one
SELECT COUNT(movies) FROM movies;
   

-- name: CreateMovie :one
INSERT INTO movies (
  title, description, director
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateMovie :exec
UPDATE movies
  set title = $1,
  description = $2,
  director= $3
WHERE movies.id = $1;

-- name: DeleteMovie :exec
DELETE FROM movies
WHERE movies.id = $1;

-- name: CreateRating :one
INSERT INTO ratings (
  movie_id, rating
) VALUES (
  $1, $2
)
RETURNING *;