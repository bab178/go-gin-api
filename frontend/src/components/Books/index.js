
const Books = ({ books }) => {
    return books.map(({ id, title, author, quantity }) => (
        <div key={`book-${id}`}>
            <p>{title} by {author}</p>
            <p>Id: {id}<br />Quantity: {quantity}</p>
            <hr />
        </div>))
};

export default Books;