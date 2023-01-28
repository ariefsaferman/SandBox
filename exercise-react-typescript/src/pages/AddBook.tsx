import InputForm from "../form/inputForm";
import { IBook } from "../interface/books";
import useFetch from "../hooks/useFetch";

export default function AddBookPage() {
  const { callRefresh } = useFetch<IBook[]>("http://localhost:3004/books");

  return (
    <div className="d-flex vh-100 flex-column align-items-center justify-content-center">
      <InputForm refresh={callRefresh}></InputForm>
    </div>
  );
}
