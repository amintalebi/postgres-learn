SELECT id, total_sales, title from 
(SELECT book_id, SUM(quantity) as total_sales from
(SELECT * from orders JOIN order_details on orders.id=order_details.order_id WHERE orders.status='SUCCESS') as sales
GROUP by book_id
order by total_sales DESC) as books_sales JOIN books on books.id=books_sales.book_id
ORDER by total_sales DESC
LIMIT 100;
