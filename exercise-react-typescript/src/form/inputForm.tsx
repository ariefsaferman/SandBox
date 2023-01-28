import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import { IBookReq } from "../interface/books";

interface Props {
  refresh: () => void;
}

function InputForm(props: Props) {
  const [title, setTitle] = useState<string>("");
  const [description, setDescription] = useState<string>("");
  const [category, setCategory] = useState<string>("Web Development");
  const [image, setImage] = useState<string>("");
  const [price, setPrice] = useState<number>(0);
  const navigate = useNavigate();

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>): void => {
    e.preventDefault();
    const request: IBookReq = {
      title: title,
      description: description,
      category: category,
      image: image,
      price: price,
    };
    const requestOptions = {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(request),
    };

    fetch("http://localhost:3004/books", requestOptions)
      .then((res) => res.json())
      .then((res) => console.log(res));

    setTitle("");
    setDescription("");
    setCategory("");
    setImage("");
    setPrice(0);
    props.refresh();
    navigate("/");
  };

  return (
    <div>
      <h1>Form Add Item</h1>

      <form onSubmit={handleSubmit}>
        <div className="input-group mb-3">
          <span className="input-group-text" id="basic-addon1">
            title
          </span>
          <input
            type="text"
            className="form-control"
            placeholder="title"
            aria-label="Username"
            aria-describedby="basic-addon1"
            onChange={(e) => {
              e.preventDefault();
              setTitle(e.target.value);
            }}
            value={title}
            required
          />
        </div>

        <div className="input-group mb-3">
          <span className="input-group-text" id="basic-addon1">
            description
          </span>
          <input
            type="text"
            className="form-control"
            placeholder="description"
            aria-label="Username"
            aria-describedby="basic-addon1"
            onChange={(e) => {
              e.preventDefault();
              setDescription(e.target.value);
            }}
            value={description}
            required
          />
        </div>

        <div className="input-group mb-3">
          <span className="input-group-text" id="basic-addon1">
            price
          </span>
          <input
            type="number"
            className="form-control"
            placeholder="0"
            aria-label="Username"
            aria-describedby="basic-addon1"
            onChange={(e) => {
              e.preventDefault();
              setPrice(parseInt(e.target.value));
            }}
            value={price}
            required
          />
        </div>

        <div className="input-group mb-3 border ">
          <span className="input-group-text" id="basic-addon1">
            category
          </span>
          <select
            className="form-select form-select-sm border border-left-0"
            aria-label=".form-select-sm example"
            onChange={(e) => {
              e.preventDefault();
              setCategory(e.target.value);
            }}
            value={category}
            required
          >
            <option value="Web Development">Web Development</option>
            <option value="Data">Data</option>
            <option value="Backend">Backend</option>
          </select>
        </div>

        <div className="input-group mb-3">
          <span className="input-group-text" id="basic-addon1">
            image
          </span>
          <input
            type="text"
            className="form-control"
            placeholder="description"
            aria-label="Username"
            aria-describedby="basic-addon1"
            onChange={(e) => {
              e.preventDefault();
              setImage(e.target.value);
            }}
            value={image}
            required
          />
        </div>

        <div className="input-group mb-3 justify-content-center">
          <button
            type="submit"
            className="btn btn-primary btn-lg w-100 text-white"
          >
            Submit
          </button>
        </div>
      </form>
    </div>
  );
}

export default InputForm;
