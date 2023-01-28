import { FormEvent } from "react";
import { useNavigate } from "react-router-dom";
import { IBook } from "../interface/books";
import useFetch from "../hooks/useFetch";
import AuthConsumer from "../hooks/useAuth";

function Home() {
  const navigate = useNavigate();
  const authState = AuthConsumer();

  const { data, callRefresh } = useFetch<IBook[]>(
    "http://localhost:3004/books"
  );

  const handleDelete = (e: FormEvent, bookId: number) => {
    e.preventDefault();
    if (window.confirm("Are you sure want to delete this item?")) {
      const requestOptions = {
        method: "DELETE",
        headers: { "Content-Type": "application/json" },
      };

      fetch("http://localhost:3004/books/" + bookId, requestOptions)
        .then((pay) => pay.json())
        .then((pay) => console.log(pay));

      callRefresh();
    }
  };

  const handleCreateClick = () => {
    navigate("/add");
  };

  const handleEditClick = (id: number) => {
    navigate(`/edit/${id}`);
  };

  const handleLogoutClick = () => {
    authState.setAuthenticated?.(false);
    localStorage.removeItem("token");
    navigate("/login");
  };

  return (
    <div className="home container">
      <div className="navbar navbar-expand-lg">
        <div className="container">
          <div className="navbar-brand fs-1 fw-bold">
            React
            <br />
            CRUD
          </div>
          <div className="navbar-nav">
            <button
              type="submit"
              className="btn btn-secondary text-white"
              onClick={handleLogoutClick}
            >
              Logout
            </button>
          </div>
        </div>
      </div>

      <div className="btn__create d-flex justify-content-start ms-3">
        <button
          type="button"
          className="btn btn-success"
          onClick={handleCreateClick}
        >
          Create
        </button>
      </div>

      <div className="table__home mt-5">
        <table className="table border border-dark border-top-0 border-end-0 border-start-0">
          <thead>
            <tr>
              <th scope="col">id</th>
              <th scope="col">Title</th>
              <th scope="col">Category</th>
              <th scope="col">Image</th>
              <th scope="col">Price</th>
              <th scope="col">Handle</th>
              <th scope="col"></th>
            </tr>
          </thead>
          <tbody>
            {data?.map((book) => {
              return (
                <tr key={book.id}>
                  <td>{book.id}</td>
                  <td>{book.title}</td>
                  <td>{book.category}</td>
                  <td>
                    <img src={book.image} alt="book" />
                  </td>
                  <td>{book.price}</td>
                  <td>
                    <button
                      type="button"
                      className="btn btn-primary text-white"
                      onClick={() => {
                        handleEditClick(book.id);
                      }}
                    >
                      Edit
                    </button>
                  </td>
                  <td>
                    <button
                      type="button"
                      className="btn btn-danger text-white"
                      onClick={(e) => handleDelete(e, book.id)}
                    >
                      Delete
                    </button>
                  </td>
                </tr>
              );
            })}
          </tbody>
        </table>
      </div>
      <div className="table"></div>
    </div>
  );
}

export default Home;
