import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { IBook, IBookReq } from "../interface/books";

interface Props {
  refresh: () => void;
  editedBook: IBook | undefined;
}

export default function EditForm(props: Props) {
  const [editTitle, setEditTitle] = useState<string>("");
  const [editDescription, setEditDescription] = useState<string>("");
  const [editCategory, setEditCategory] = useState<string>("Web Development");
  const [editImage, setEditImage] = useState<string>("");
  const [editPrice, setEditPrice] = useState<number>(0);
  const navigate = useNavigate();

  const handleEditSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const request: IBookReq = {
      title: editTitle,
      description: editDescription,
      category: editCategory,
      image: editImage,
      price: editPrice,
    };

    const requestOptions = {
      method: "PATCH",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(request),
    };

    fetch("http://localhost:3004/books/" + props.editedBook?.id, requestOptions)
      .then((props) => props.json())
      .then((props) => console.log(props));

    setEditTitle("");
    setEditDescription("");
    setEditCategory("");
    setEditImage("");
    setEditPrice(0);
    props.refresh();
    navigate("/");
  };

  useEffect(() => {
    if (props.editedBook) {
      setEditTitle(props.editedBook.title);
      setEditDescription(props.editedBook.description);
      setEditCategory(props.editedBook.category);
      setEditImage(props.editedBook.image);
      setEditPrice(props.editedBook.price);
    }
  }, [props.editedBook]);

  return (
    <div>
      <h2>Form Edit Item</h2>

      <form onSubmit={(e) => handleEditSubmit(e)}>
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
              setEditTitle(e.target.value);
            }}
            value={editTitle}
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
              setEditDescription(e.target.value);
            }}
            value={editDescription}
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
              setEditPrice(parseInt(e.target.value));
            }}
            value={editPrice}
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
              setEditCategory(e.target.value);
            }}
            value={editCategory}
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
              setEditImage(e.target.value);
            }}
            value={editImage}
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
