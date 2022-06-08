import React, { useRef } from 'react';
import './CreateBook.css'

const CreateGroup = React.forwardRef(({ name }, ref) => {
    return (
        <div className='group'>
            <label>{name}: </label>
            <input name={name.toLowerCase()} ref={ref} />
        </div>);
})

const CreateBook = ({ createNewBook }) => {

    const idRef = useRef({});
    const titleRef = useRef({});
    const authorRef = useRef({});
    const quantityRef = useRef({});


    const onClick = () => {
        createNewBook({
            Id: idRef.current.value,
            Title: titleRef.current.value,
            Author: authorRef.current.value,
            Quantity: parseInt(quantityRef.current.value)
        })
    }

    return (
        <div>
            <br />
            <CreateGroup name="Id" ref={idRef} />
            <CreateGroup name="Title" ref={titleRef} />
            <CreateGroup name="Author" ref={authorRef} />
            <CreateGroup name="Quantity" ref={quantityRef} />
            <button type="submit" onClick={onClick}>Create New Book</button>
        </div>
    )
};

export default CreateBook;