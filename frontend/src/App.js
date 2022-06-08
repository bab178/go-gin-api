import './App.css';
import { useEffect, useState, useRef } from 'react';
import Books from './components/Books';
import CreateBook from './components/CreateBook';

const host = "http://localhost:8080";

const App = () => {

  const selectRef = useRef({});
  const [books, setBooks] = useState([]);

  useEffect(() => {
    const fetchBooks = () => {
      fetch(`${host}/books`)
        .then(res => res.json())
        .then(data => setBooks(data))
    };

    fetchBooks();
  }, []);

  const postData = (url = '', data = {}) => {
    return fetch(url, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(data)
    });
  }

  const returnBook = () => {
    const id = selectRef.current.value
    postData(`${host}/books/return/${id}`)
      .then(res => res.json())
      .then(data => { data.length > 0 ? setBooks(data) : alert(data.message) });
  }

  const checkoutBook = () => {
    const id = selectRef.current.value
    postData(`${host}/books/checkout/${id}`)
      .then(res => res.json())
      .then(data => { data.length > 0 ? setBooks(data) : alert(data.message) });
  };

  const createNewBook = (bookData) => {
    postData(`${host}/books`, { ...bookData })
      .then(res => res.json())
      .then(data => { data.length > 0 ? setBooks(data) : alert(data.message) });
  };

  return (
    <div className="App">
      {books.length > 0 ? (
        <>
          <Books books={books} />
          <select ref={selectRef}>
            {books.map(({ title, id }) => (
              <option key={`option-${id}`} value={id}>{title}</option>
            ))}
          </select>
          <br />
          <button onClick={returnBook}>Return 1 copy</button>
          <button onClick={checkoutBook}>Checkout 1 copy</button>
          <CreateBook createNewBook={createNewBook} />
        </>
      ) : (
        <h1>Loading...</h1>
      )}
    </div>
  );
}

export default App;
