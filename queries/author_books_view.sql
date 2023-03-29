WITH
    book_rating as (SELECT book_id, avg(rate) as rating from book_reviews GROUP by book_id)
    SELECT author, avg(rating) as author_rating from books join book_rating on books.id=book_rating.book_id group by author order by author_rating DESC;