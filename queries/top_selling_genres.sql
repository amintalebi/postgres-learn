WITH sales as (SELECT * from orders JOIN order_details on orders.id=order_details.order_id WHERE orders.status='SUCCESS'),
books_sales as (SELECT book_id, SUM(quantity) as total_sales from sales GROUP by book_id)
SELECT categories.name, categories.description, sum(total_sales) as genre_sales from categories join 
(SELECT id,title,total_sales,category_id from books_sales
JOIN books on books.id=books_sales.book_id) as top_selling_books
on categories.id=top_selling_books.category_id
GROUP by category_id, categories.name, categories.description
order by genre_sales DESC;
